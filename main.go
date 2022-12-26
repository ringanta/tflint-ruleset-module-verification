package main

import (
	"github.com/ringanta/tflint-ruleset-module-signature/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "module-signature",
			Version: "0.1.0",
			Rules:   rules.Rules,
		},
	})
}
