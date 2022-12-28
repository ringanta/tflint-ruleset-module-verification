# module_verification_nonregistry_source

Explicitly allows non Terraform Registry modules.

## Configuration

Name | Description | Default | Type
--- | --- | --- | ---
allowed_modules | List of allowed modules prefix | [] | List of string

```hcl
rule "module_verification_nonregistry_source" {
  enabled = true
  allowed_modules = [] # default
}
```

## Example

```
tflint
2 issue(s) found:

Error: module "local_fail" source is not on the allowed modules list (module_verification_nonregistry_source)

  on main.tf line 1:
   1: module "local_fail" {

Reference: https://github.com/ringanta/tflint-ruleset-module-verification/blob/v0.1.0/docs/rules/module_verification_nonregistry_source.md
```

## Why

Module is external code that needs to be vet before being used. Explicitly allow module usage is a good practice.

## How to fix

Use the following TFLint config to explicitly allow module from local source
```
rule "module_verification_nonregistry_source" {
	enabled = true

    allowed_modules = [
        "../.."
    ]
}
```
