package rules

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-getter"
	tfaddr "github.com/hashicorp/terraform-registry-address"
	moduleverification "github.com/ringanta/tflint-ruleset-module-verification/module-verification"
	"github.com/ringanta/tflint-ruleset-module-verification/project"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// ModuleVerificationLocalSourceRule checks local module against certain rules
type ModuleVerificationLocalSourceRule struct {
	tflint.DefaultRule
}

// ModuleVerificationLocalSourceRuleConfig is the config structure for the ModuleSignatureLocalSourceRule rule
type ModuleVerificationLocalSourceRuleConfig struct {
	AllowList []string `hclext:"allowed_modules,optional"`
}

// NewModuleVerificationLocalSourceRule returns new rule with default attributes
func NewModuleVerificationLocalSourceRule() *ModuleVerificationLocalSourceRule {
	return &ModuleVerificationLocalSourceRule{}
}

// Name returns the rule name
func (r *ModuleVerificationLocalSourceRule) Name() string {
	return "module_verification_local_source"
}

// Enabled returns whether the rule enabled by default
func (r *ModuleVerificationLocalSourceRule) Enabled() bool {
	return true
}

// Severity returns severity of the rule
func (r *ModuleVerificationLocalSourceRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *ModuleVerificationLocalSourceRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks if module source is local
func (r *ModuleVerificationLocalSourceRule) Check(rr tflint.Runner) error {
	runner := moduleverification.NewRunner(rr)

	path, err := runner.GetModulePath()
	if err != nil {
		return err
	}
	if !path.IsRoot() {
		// This rule doesn't evaluate child module
		return nil
	}

	config := ModuleVerificationLocalSourceRuleConfig{AllowList: make([]string, 0)}
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

func (r *ModuleVerificationLocalSourceRule) checkModule(runner tflint.Runner, module *moduleverification.ModuleCall, config ModuleVerificationLocalSourceRuleConfig) error {
	_, err := tfaddr.ParseModuleSource(module.Source)
	if err != nil {
		source, err := getter.Detect(module.Source, filepath.Dir(module.DefRange.Filename), getter.Detectors)
		if err != nil {
			return err
		}

		u, err := url.Parse(source)
		if err != nil {
			return err
		}

		if u.Scheme == "file" {
			for _, item := range config.AllowList {
				if strings.HasPrefix(module.Source, item) {
					return nil
				}
			}

			return runner.EmitIssue(
				r,
				fmt.Sprintf("module %q should not use local source", module.Name),
				module.DefRange,
			)
		}
	}

	return nil
}
