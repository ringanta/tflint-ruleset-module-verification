package rules

import (
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/hashicorp/go-getter"
	tfaddr "github.com/hashicorp/terraform-registry-address"
	modulesignature "github.com/ringanta/tflint-ruleset-module-signature/module-signature"
	"github.com/ringanta/tflint-ruleset-module-signature/project"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// ModuleSignatureLocalSourceRule checks local module against certain rules
type ModuleSignatureLocalSourceRule struct {
	tflint.DefaultRule
}

// ModuleSignatureLocalSourceRuleConfig is the config structure for the ModuleSignatureLocalSourceRule rule
type ModuleSignatureLocalSourceRuleConfig struct {
	Allow bool `hclext:"allow,optional"`
}

// NewModuleSignatureLocalSourceRule returns new rule with default attributes
func NewModuleSignatureLocalSourceRule() *ModuleSignatureLocalSourceRule {
	return &ModuleSignatureLocalSourceRule{}
}

// Name returns the rule name
func (r *ModuleSignatureLocalSourceRule) Name() string {
	return "module_signature_local_source"
}

// Enabled returns whether the rule enabled by default
func (r *ModuleSignatureLocalSourceRule) Enabled() bool {
	return true
}

// Severity returns severity of the rule
func (r *ModuleSignatureLocalSourceRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *ModuleSignatureLocalSourceRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks if module source is local
func (r *ModuleSignatureLocalSourceRule) Check(rr tflint.Runner) error {
	runner := modulesignature.NewRunner(rr)

	path, err := runner.GetModulePath()
	if err != nil {
		return err
	}
	if !path.IsRoot() {
		// This rule doesn't evaluate child module
		return nil
	}

	config := ModuleSignatureLocalSourceRuleConfig{Allow: true}
	if err := runner.DecodeRuleConfig(r.Name(), &config); err != nil {
		return err
	}

	calls, diags := runner.GetModuleCalls()
	if diags.HasErrors() {
		return diags
	}

	for _, call := range calls {
		if err := r.checkModule(runner, call, config); err != nil {
			return err
		}
	}

	return nil
}

func (r *ModuleSignatureLocalSourceRule) checkModule(runner tflint.Runner, module *modulesignature.ModuleCall, config ModuleSignatureLocalSourceRuleConfig) error {
	_, err := tfaddr.ParseModuleSource(module.Source)
	if err != nil {
		source, err := getter.Detect(module.Source, filepath.Dir(module.DefRange.Filename), getter.Detectors)
		if err != nil {
			return err
		}
		logger.Debug(fmt.Sprintf("Source %s", source))

		u, err := url.Parse(source)
		if err != nil {
			return err
		}

		if u.Scheme == "file" && !config.Allow {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("module %q should not use local source", module.Name),
				module.DefRange,
			)
		}
	}

	return nil
}
