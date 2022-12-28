package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestModuleVerificationRegistrySource(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Config   string
		Expected helper.Issues
	}{
		{
			Name:    "deny",
			Content: testModuleVerificationRegistrySource,
			Expected: helper.Issues{
				{
					Rule:    NewModuleVerificationRegistrySourceRule(),
					Message: `module "registry" is not in the list of allowed modules from Terraform Registry`,
					Range: hcl.Range{
						Filename: "modules.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 18},
					},
				},
			},
		},
		{
			Name:    "allow_list",
			Content: testModuleVerificationRegistrySource,
			Config: `
rule "module_signature_registry_source" {
	enabled = true
	allowed_module {
		source = "terraform-aws-modules/vpc"
		versions = [
			"~> 3.18"
		]
	}
}			
`,
			Expected: helper.Issues{},
		},
		{
			Name:    "allow_deny_list",
			Content: testModuleVerificationRegistrySource,
			Config: `
rule "module_signature_registry_source" {
	enabled = true
	allowed_module {
		source = "terraform-aws-modules/vpc"
		versions = [
			"~> 3.18"
		]
	}

	denied_module {
		source = "terraform-aws-modules/vpc"
		versions = [
			"3.18.1"
		]
	}
}			
`,
			Expected: helper.Issues{
				{
					Rule:    NewModuleVerificationRegistrySourceRule(),
					Message: `module "registry" version "3.18.1" from Terraform Registry is in the deny list`,
					Range: hcl.Range{
						Filename: "modules.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 18},
					},
				},
			},
		},
		{
			Name:    "multiple_allow",
			Content: testModuleVerificationRegistrySource,
			Config: `
rule "module_signature_registry_source" {
	enabled = true

	allowed_module {
		source = "terraform-aws-modules/ec2"
		versions = [
			"~> 3.18"
		]
	}

	allowed_module {
		source = "terraform-aws-modules/vpc"
		versions = [
			"~> 3.18"
		]
	}
}
`,
			Expected: helper.Issues{},
		},
		{
			Name:    "multiple_deny",
			Content: testModuleVerificationRegistrySource,
			Config: `
rule "module_signature_registry_source" {
	enabled = true
	allowed_module {
		source = "terraform-aws-modules/vpc"
		versions = [
			"~> 3.18"
		]
	}

	denied_module {
		source = "terraform-aws-modules/ec2"
		versions = [
			"~> 3.18"
		]
	}

	denied_module {
		source = "terraform-aws-modules/vpc"
		versions = [
			"3.18.1"
		]
	}
}			
`,
			Expected: helper.Issues{
				{
					Rule:    NewModuleVerificationRegistrySourceRule(),
					Message: `module "registry" version "3.18.1" from Terraform Registry is in the deny list`,
					Range: hcl.Range{
						Filename: "modules.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 18},
					},
				},
			},
		},
	}

	rule := NewModuleVerificationRegistrySourceRule()
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			runner := testRunner(t, map[string]string{"modules.tf": tc.Content, ".tflint.hcl": tc.Config})

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, tc.Expected, runner.Runner.(*helper.Runner).Issues)
		})
	}
}

const testModuleVerificationRegistrySource = `
module "registry" {
	source = "terraform-aws-modules/vpc/aws"
	version = "3.18.1"
}
`
