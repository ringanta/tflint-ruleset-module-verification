plugin "terraform" {
  enabled = false
}

plugin "module-verification" {
  enabled = true
}

rule "module_verification_registry_source" {
	enabled = true

  // List of allowed modules
  allowed_module {
    source = "ringanta/backend"
    versions = [
      "~> 1.0"
    ]
  }
}
