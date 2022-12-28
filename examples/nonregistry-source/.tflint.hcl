plugin "terraform" {
  enabled = false
}

plugin "module-verification" {
  enabled = true
}

rule "module_verification_nonregistry_source" {
	enabled = true

  // List of allowed module prefix
  allowed_modules = [
    "../../terraform-modules",
    "github.com/example/",
    "git@github.com:example/"
  ]
}
