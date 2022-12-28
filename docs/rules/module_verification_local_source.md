# module_verification_local_source

Explicitly allow locally sourced module.

## Configuration

Name | Description | Default | Type
--- | --- | --- | ---
allowed_modules | List of allowed modules prefix | [] | List of string

```hcl
rule "terraform_module_version" {
  enabled = true
  allowed_modules = [] # default
}
```

## Example

```
tflint
1 issue(s) found:

Error: module "local_fail" should not use local source (module_signature_local_source)

  on main.tf line 1:
   1: module "local_fail" {

Reference: https://github.com/ringanta/tflint-ruleset-module-signature/blob/v0.1.0/docs/rules/module_signature_local_source.md
```

## Why

Module is external code that needs to be vet before being used. Explicitly allow module usage is a good practice.

## How to fix

Use the following TFLint config to explicitly allow module from local source
```
rule "module_signature_local_source" {
	enabled = true

    allowed_modules = [
        "../.."
    ]
}
```
