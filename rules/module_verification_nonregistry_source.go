package rules

import (
	"fmt"
	"strings"

	tfaddr "github.com/hashicorp/terraform-registry-address"
	moduleverification "github.com/ringanta/tflint-ruleset-module-verification/module-verification"
	"github.com/ringanta/tflint-ruleset-module-verification/project"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// ModuleVerificationNonRegistrySourceRule checks local module against certain rules
type ModuleVerificationNonRegistrySourceRule struct {
	tflint.DefaultRule
}

// ModuleVerificationNonRegistrySourceRuleConfig is the config structure for the ModuleSignatureLocalSourceRule rule
type ModuleVerificationNonRegistrySourceRuleConfig struct {
	AllowList []string `hclext:"allowed_modules,optional"`
}

// NewModuleVerificationNonRegistrySourceRule returns new rule with default attributes
func NewModuleVerificationNonRegistrySourceRule() *ModuleVerificationNonRegistrySourceRule {
	return &ModuleVerificationNonRegistrySourceRule{}
}

// Name returns the rule name
func (r *ModuleVerificationNonRegistrySourceRule) Name() string {
	return "module_verification_nonregistry_source"
}

// Enabled returns whether the rule enabled by default
func (r *ModuleVerificationNonRegistrySourceRule) Enabled() bool {
	return true
}

// Severity returns severity of the rule
func (r *ModuleVerificationNonRegistrySourceRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *ModuleVerificationNonRegistrySourceRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks if module source is local
func (r *ModuleVerificationNonRegistrySourceRule) Check(rr tflint.Runner) error {
	runner := moduleverification.NewRunner(rr)

	path, err := runner.GetModulePath()
	if err != nil {
		return err
	}
	if !path.IsRoot() {
		// This rule doesn't evaluate child module
		return nil
	}

	config := ModuleVerificationNonRegistrySourceRuleConfig{AllowList: make([]string, 0)}
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

func (r *ModuleVerificationNonRegistrySourceRule) checkModule(runner tflint.Runner, module *moduleverification.ModuleCall, config ModuleVerificationNonRegistrySourceRuleConfig) error {
	_, err := tfaddr.ParseModuleSource(module.Source)
	if err != nil {
		for _, item := range config.AllowList {
			if strings.HasPrefix(module.Source, item) {
				return nil
			}
		}

		return runner.EmitIssue(
			r,
			fmt.Sprintf("module %q source is not on the allowed modules list", module.Name),
			module.DefRange,
		)
	}

	return nil
}
