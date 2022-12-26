package modulesignature

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// Runner is a custom runner that provides helper function for this ruleset
type Runner struct {
	tflint.Runner
}

// NewRunner returns a new custom runner.
func NewRunner(runner tflint.Runner) *Runner {
	return &Runner{Runner: runner}
}

// GetModuleCalls returns all "module" blocks, including uncreated module calls.
func (r *Runner) GetModuleCalls() ([]*ModuleCall, hcl.Diagnostics) {
	calls := []*ModuleCall{}
	diags := hcl.Diagnostics{}

	body, err := r.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "module",
				LabelNames: []string{"name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "source"},
						{Name: "version"},
					},
				},
			},
		},
	}, &tflint.GetModuleContentOption{ExpandMode: tflint.ExpandModeNone})
	if err != nil {
		return calls, hcl.Diagnostics{
			{
				Severity: hcl.DiagError,
				Summary:  "failed to call GetModuleContent()",
				Detail:   err.Error(),
			},
		}
	}

	for _, block := range body.Blocks {
		call, decodeDiags := decodeModuleCall(block)
		diags = diags.Extend(decodeDiags)
		if decodeDiags.HasErrors() {
			continue
		}
		calls = append(calls, call)
	}

	return calls, diags
}
