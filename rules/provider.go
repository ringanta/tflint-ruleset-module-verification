package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

var moduleSignatureRules = []tflint.Rule{}

// Rules is a list of all rules
var Rules []tflint.Rule

func init() {
	Rules = append(Rules, moduleSignatureRules...)
}
