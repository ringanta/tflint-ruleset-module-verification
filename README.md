# TFLint Module Signature Ruleset
[![Build Status](https://github.com/ringanta/tflint-ruleset-module-signature/workflows/build/badge.svg?branch=main)](https://github.com/ringanta/tflint-ruleset-module-signature/actions)

TFlint plugin to validate module sources and signature.

## Requirements

- TFLint v0.40+
- Go v1.19

## Installation

You can install the plugin with `tflint --init`. Declare a config in `.tflint.hcl` as follows:

```hcl
plugin "module-signature" {
  enabled = true

  version = "0.1.0"
  source  = "github.com/ringanta/tflint-ruleset-module-signature"

  signing_key = <<-KEY

  KEY
}
```

## Rules

|Name|Description|Severity|Enabled|Link|
| --- | --- | --- | --- | --- |


## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```

You can run the built plugin like the following:

```
$ cat << EOS > .tflint.hcl
plugin "module-signature" {
  enabled = true
}
EOS
$ tflint
```

## Directory layout

Here is layout of directory of the repository

- `docs` contains documentation of module-signature ruleset
- `examples` contains list of examples how the rules can be used
