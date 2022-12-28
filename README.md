# TFLint Module Verification Ruleset
[![Build Status](https://github.com/ringanta/tflint-ruleset-module-verification/workflows/build/badge.svg?branch=main)](https://github.com/ringanta/tflint-ruleset-module-verification/actions)

TFlint plugin to validate modules source.
The plugin contains ruleset to validate both Terraform Registry and non Terraform Registry modules.

## Requirements

- TFLint v0.40+
- Go v1.19

## Installation

You can install the plugin with `tflint --init`. Declare a config in `.tflint.hcl` as follows:

```hcl
plugin "module-verification" {
  enabled = true

  version = "0.1.0"
  source  = "github.com/ringanta/tflint-ruleset-module-verification"

  signing_key = <<-KEY

  KEY
}
```

## Rules

See [Rules](./docs/rules/README.md)

## Building the plugin

Clone the repository locally and run the following command:

```
make
```

You can run the test suites with the following command:
```shell
make test
```

You can easily install the built plugin with the following:

```
make install
```

You can run the built plugin like the following:

```
$ cat << EOS > .tflint.hcl
plugin "module-verification" {
  enabled = true
}
EOS
$ tflint
```

## Directory layout

Here is layout of directory of the repository

- `docs` contains documentation of module-verification ruleset
- `examples` contains list of examples how the rules can be used
- `rules` contains ruleset implementation of the module-verification plugin
