package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestModuleVerificationNonRegistrySource(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Config   string
		Expected helper.Issues
	}{
		{
			Name: "deny_local",
			Content: `
module "local" {
	source = "./example"
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewModuleVerificationNonRegistrySourceRule(),
					Message: `module "local" source is not on the allowed modules list`,
					Range: hcl.Range{
						Filename: "modules.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 15},
					},
				},
			},
		},
		{
			Name: "allow_local",
			Content: `
module "local" {
	source = "./example"
}
`,
			Config: `
rule "module_verification_nonregistry_source" {
	enabled = true
	
	allowed_modules = [
		"./example",
	]
}			
`,
			Expected: helper.Issues{},
		},
		{
			Name: "deny_https",
			Content: `
module "https" {
	source = "https://example.com/example-module.zip"
}	
`,
			Expected: helper.Issues{
				{
					Rule:    NewModuleVerificationNonRegistrySourceRule(),
					Message: `module "https" source is not on the allowed modules list`,
					Range: hcl.Range{
						Filename: "modules.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 15},
					},
				},
			},
		},
		{
			Name: "allow_https",
			Content: `
module "https" {
	source = "https://example.com/example-module.zip"
}	
`,
			Config: `
rule "module_verification_nonregistry_source" {
	enabled = true
	
	allowed_modules = [
		"https://example.com/example-module.zip",
	]
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "deny_github",
			Content: `
module "github" {
	source = "github.com/example/example-module"
}	
`,
			Expected: helper.Issues{
				{
					Rule:    NewModuleVerificationNonRegistrySourceRule(),
					Message: `module "github" source is not on the allowed modules list`,
					Range: hcl.Range{
						Filename: "modules.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 16},
					},
				},
			},
		},
		{
			Name: "allow_github",
			Content: `
module "github" {
	source = "github.com/example/example-module"
}	
`,
			Config: `
rule "module_verification_nonregistry_source" {
	enabled = true
	
	allowed_modules = [
		"github.com/example/",
	]
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "deny_bitbucket",
			Content: `
module "bitbucket" {
	source = "bitbucket.org/example/example-module"
}	
`,
			Expected: helper.Issues{
				{
					Rule:    NewModuleVerificationNonRegistrySourceRule(),
					Message: `module "bitbucket" source is not on the allowed modules list`,
					Range: hcl.Range{
						Filename: "modules.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 19},
					},
				},
			},
		},
		{
			Name: "allow_bitbucket",
			Content: `
module "bitbucket" {
	source = "bitbucket.org/example/example-module"
}	
`,
			Config: `
rule "module_verification_nonregistry_source" {
	enabled = true
	
	allowed_modules = [
		"bitbucket.org/example/",
	]
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "deny_generic_git",
			Content: `
module "generic_git" {
	source = "git::https://example.com/example-module.git"
}	
`,
			Expected: helper.Issues{
				{
					Rule:    NewModuleVerificationNonRegistrySourceRule(),
					Message: `module "generic_git" source is not on the allowed modules list`,
					Range: hcl.Range{
						Filename: "modules.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 21},
					},
				},
			},
		},
		{
			Name: "allow_generic_git",
			Content: `
module "generic_git" {
	source = "git::https://example.com/example-module.git"
}	
`,
			Config: `
rule "module_verification_nonregistry_source" {
	enabled = true
	
	allowed_modules = [
		"git::https://example.com/",
	]
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "deny_generic_hg",
			Content: `
module "generic_hg" {
	source = "hg::https://example.com/example-module.hg"
}	
`,
			Expected: helper.Issues{
				{
					Rule:    NewModuleVerificationNonRegistrySourceRule(),
					Message: `module "generic_hg" source is not on the allowed modules list`,
					Range: hcl.Range{
						Filename: "modules.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 20},
					},
				},
			},
		},
		{
			Name: "allow_generic_hg",
			Content: `
module "generic_hg" {
	source = "hg::https://example.com/example-module.hg"
}	
`,
			Config: `
rule "module_verification_nonregistry_source" {
	enabled = true
	
	allowed_modules = [
		"hg::https://example.com/",
	]
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "deny_s3",
			Content: `
module "s3" {
	source = "s3::https://example-modules.s3.ap-southeast-1.amazonaws.com/example-module.hg"
}	
`,
			Expected: helper.Issues{
				{
					Rule:    NewModuleVerificationNonRegistrySourceRule(),
					Message: `module "s3" source is not on the allowed modules list`,
					Range: hcl.Range{
						Filename: "modules.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 12},
					},
				},
			},
		},
		{
			Name: "allow_s3",
			Content: `
module "s3" {
	source = "s3::https://example-modules.s3.ap-southeast-1.amazonaws.com/example-module.hg"
}	
`,
			Config: `
rule "module_verification_nonregistry_source" {
	enabled = true
	
	allowed_modules = [
		"s3::https://example-modules.s3.ap-southeast-1.amazonaws.com/",
	]
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "deny_gcs",
			Content: `
module "gcs" {
	source = "gcs::https://www.googleapis.com/storage/v1/example/example-module.zip"
}	
`,
			Expected: helper.Issues{
				{
					Rule:    NewModuleVerificationNonRegistrySourceRule(),
					Message: `module "gcs" source is not on the allowed modules list`,
					Range: hcl.Range{
						Filename: "modules.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 13},
					},
				},
			},
		},
		{
			Name: "allow_gcs",
			Content: `
module "gcs" {
	source = "gcs::https://www.googleapis.com/storage/v1/example/example-module.zip"
}	
`,
			Config: `
rule "module_verification_nonregistry_source" {
	enabled = true
	
	allowed_modules = [
		"gcs::https://www.googleapis.com/storage/v1/example/",
	]
}
`,
			Expected: helper.Issues{},
		},
	}

	rule := NewModuleVerificationNonRegistrySourceRule()
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
