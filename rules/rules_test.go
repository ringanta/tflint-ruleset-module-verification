package rules

import (
	"testing"

	modulesignature "github.com/ringanta/tflint-ruleset-module-signature/module-signature"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func testRunner(t *testing.T, files map[string]string) *modulesignature.Runner {
	return modulesignature.NewRunner(helper.TestRunner(t, files))
}
