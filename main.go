package main

import (
	moduleverification "github.com/ringanta/tflint-ruleset-module-verification/module-verification"
	"github.com/ringanta/tflint-ruleset-module-verification/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &moduleverification.RuleSet{
			BuiltinRuleSet: tflint.BuiltinRuleSet{
				Name:    "module-verification",
				Version: "0.1.0",
			},
			PresetRules: rules.PresetRules,
		},
	})
}
