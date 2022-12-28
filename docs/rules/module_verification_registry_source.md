# module_verification_registry_source

Explicitly allows module from Terraform Registry.

## Configuration

Name | Description | Default | Type
--- | --- | --- | ---
allowed_module | Module that's allowed to be used from Terraform Tegistry | `{}` | Block with `source` and `versions` attributes
denied_module | Module that's explicitly denied to be used from Terraform Registry | `{}` | Block with `source` and `versions` attributes

```hcl
rule "module_verification_registry_source" {
  enabled = true
  allowed_module {
    source = "ringanta/backend/aws"
    versions = [
        "~> 1.0"
    ]
  }
}
```

## Example

```
tflint
1 issue(s) found:

Error: module "registry_fail" is not in the list of allowed modules from Terraform Registry (module_verification_registry_source)

  on main.tf line 1:
   1: module "registry_fail" {

Reference: https://github.com/ringanta/tflint-ruleset-module-verification/blob/v0.1.0/docs/rules/module_verification_registry_source.md
```

## Why

Consumption of module from Terraform Registry needs to be explicitly allowed.

## How to fix

Use the following TFLint config to explicitly allow module from local source
```
rule "module_verification_registry_source" {
  enabled = true

  allowed_module {
    source = "ringanta/backend"
    versions = [
      "~> 1.0"
    ]
  }
}
```
