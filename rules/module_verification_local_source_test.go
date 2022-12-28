package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestModuleVerificationLocalSource(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Config   string
		Expected helper.Issues
	}{
		{
			Name: "local_deny",
			Content: `
module "local" {
	source = "../.."
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewModuleVerificationLocalSourceRule(),
					Message: `module "local" should not use local source`,
					Range: hcl.Range{
						Filename: "modules.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 15},
					},
				},
			},
		},
		{
			Name: "local_allow",
			Content: `
module "local" {
	source = "../.."
}
`,
			Config: `
rule "module_signature_local_source" {
	enabled = true
	allow = true
}			
`,
			Expected: helper.Issues{},
		},
	}

	rule := NewModuleVerificationLocalSourceRule()
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
