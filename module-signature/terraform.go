package modulesignature

import (
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
)

// ModuleCall represents a "module" block.
type ModuleCall struct {
	Name        string
	DefRange    hcl.Range
	Source      string
	SourceAttr  *hclext.Attribute
	Version     version.Constraints
	VersionAttr *hclext.Attribute
}

// @see https://github.com/hashicorp/terraform/blob/v1.2.7/internal/configs/module_call.go#L36-L224
func decodeModuleCall(block *hclext.Block) (*ModuleCall, hcl.Diagnostics) {
	module := &ModuleCall{
		Name:     block.Labels[0],
		DefRange: block.DefRange,
	}
	diags := hcl.Diagnostics{}

	if source, exists := block.Body.Attributes["source"]; exists {
		module.SourceAttr = source
		sourceDiags := gohcl.DecodeExpression(source.Expr, nil, &module.Source)
		diags = diags.Extend(sourceDiags)
	}

	if versionAttr, exists := block.Body.Attributes["version"]; exists {
		module.VersionAttr = versionAttr

		var versionVal string
		versionDiags := gohcl.DecodeExpression(versionAttr.Expr, nil, &versionVal)
		diags = diags.Extend(versionDiags)
		if diags.HasErrors() {
			return module, diags
		}

		constraints, err := version.NewConstraint(versionVal)
		if err != nil {
			diags = append(diags, &hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  "Invalid version constraint",
				Detail:   "This string does not use correct version constraint syntax.",
				Subject:  versionAttr.Expr.Range().Ptr(),
			})
		}
		module.Version = constraints
	}

	return module, diags
}
