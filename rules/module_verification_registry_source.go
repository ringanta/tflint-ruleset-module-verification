package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/go-version"
	tfaddr "github.com/hashicorp/terraform-registry-address"
	modulesignature "github.com/ringanta/tflint-ruleset-module-signature/module-signature"
	"github.com/ringanta/tflint-ruleset-module-signature/project"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// See https://github.com/terraform-linters/tflint-ruleset-terraform/blob/ed5566cffeb23892e3e0a3a9bb890972fe2ab1ba/rules/terraform_module_version.go#L15
var exactVersionRegexp = regexp.MustCompile(`^=?\s*` + `(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

// ModuleVerificationRegistrySourceRule checks local module against certain rules
type ModuleVerificationRegistrySourceRule struct {
	tflint.DefaultRule
}

// ListItemConfig represents entry on the allow list and the deny list
type ListItemConfig struct {
	Source   string   `hclext:"source,optional"`
	Versions []string `hclext:"versions,optional"`
}

// ModuleVerificationRegistrySourceRuleConfig is the config structure for the ModuleSignatureRegistrySourceRule rule
type ModuleVerificationRegistrySourceRuleConfig struct {
	AllowList []*ListItemConfig `hclext:"allowed_module,block"`
	DenyList  []*ListItemConfig `hclext:"denied_module,block"`
}

// NewModuleVerificationRegistrySourceRule returns new rule with default attributes
func NewModuleVerificationRegistrySourceRule() *ModuleVerificationRegistrySourceRule {
	return &ModuleVerificationRegistrySourceRule{}
}

// Name returns the rule name
func (r *ModuleVerificationRegistrySourceRule) Name() string {
	return "module_signature_registry_source"
}

// Enabled returns whether the rule enabled by default
func (r *ModuleVerificationRegistrySourceRule) Enabled() bool {
	return true
}

// Severity returns severity of the rule
func (r *ModuleVerificationRegistrySourceRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *ModuleVerificationRegistrySourceRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks if module source is Registry and validate the source against defined rules
func (r *ModuleVerificationRegistrySourceRule) Check(rr tflint.Runner) error {
	runner := modulesignature.NewRunner(rr)

	path, err := runner.GetModulePath()
	if err != nil {
		return err
	}
	if !path.IsRoot() {
		// This rule doesn't evaluate child module
		return nil
	}

	config := ModuleVerificationRegistrySourceRuleConfig{AllowList: make([]*ListItemConfig, 0), DenyList: make([]*ListItemConfig, 0)}
	if err := runner.DecodeRuleConfig(r.Name(), &config); err != nil {
		return err
	}

	logger.Debug(fmt.Sprintf("Allow list %v, Deny list %v", config.AllowList, config.DenyList))

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

func (r *ModuleVerificationRegistrySourceRule) checkModule(runner tflint.Runner, module *modulesignature.ModuleCall, config ModuleVerificationRegistrySourceRuleConfig) error {
	_, err := tfaddr.ParseModuleSource(module.Source)
	if err != nil {
		return nil
	}

	logger.Debug(fmt.Sprintf("Module %q source %q is coming from Terraform Registry", module.Name, module.Source))

	if len(module.Version) > 1 {
		return runner.EmitIssue(
			r,
			fmt.Sprintf("module %q should specify an exact version, but multiple constraints were found", module.Name),
			module.VersionAttr.Range,
		)
	}

	moduleVersion := module.Version[0].String()
	if !exactVersionRegexp.MatchString(moduleVersion) {
		return runner.EmitIssue(
			r,
			fmt.Sprintf("module %q should specify an exact version, but a range was found", module.Name),
			module.VersionAttr.Range,
		)
	}

	// Deny module usage if if matches any of entry in the deny list
	for _, item := range config.DenyList {
		if strings.HasPrefix(module.Source, item.Source) {
			for _, v := range item.Versions {
				denyConstraint, err := version.NewConstraint(v)
				if err != nil {
					return err
				}

				moduleVersionObj, err := version.NewVersion(moduleVersion)
				if err != nil {
					return err
				}

				if denyConstraint.Check(moduleVersionObj) {
					return runner.EmitIssue(
						r,
						fmt.Sprintf("module %q version %q from Terraform Registry is in the deny list", module.Name, moduleVersion),
						module.DefRange,
					)
				}
			}
		}
	}

	// Allow module usage if they are in the allow list
	for _, item := range config.AllowList {
		if strings.HasPrefix(module.Source, item.Source) {
			for _, v := range item.Versions {
				allowConstraint, err := version.NewConstraint(v)
				if err != nil {
					return err
				}

				moduleVersionObj, err := version.NewVersion(moduleVersion)
				if err != nil {
					return err
				}

				if allowConstraint.Check(moduleVersionObj) {
					return nil
				}
			}
		}
	}

	// Deny module usage by default
	return runner.EmitIssue(
		r,
		fmt.Sprintf("module %q is not in the list of allowed modules from Terraform Registry", module.Name),
		module.DefRange,
	)
}
