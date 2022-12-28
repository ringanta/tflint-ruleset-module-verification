package main

import (
	modulesignature "github.com/ringanta/tflint-ruleset-module-signature/module-signature"
	"github.com/ringanta/tflint-ruleset-module-signature/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &modulesignature.RuleSet{
			BuiltinRuleSet: tflint.BuiltinRuleSet{
				Name:    "module-verification",
				Version: "0.1.0",
			},
			PresetRules: rules.PresetRules,
		},
	})
}
