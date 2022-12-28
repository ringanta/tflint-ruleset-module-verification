package rules

import (
	"testing"

	moduleverification "github.com/ringanta/tflint-ruleset-module-verification/module-verification"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func testRunner(t *testing.T, files map[string]string) *moduleverification.Runner {
	return moduleverification.NewRunner(helper.TestRunner(t, files))
}
