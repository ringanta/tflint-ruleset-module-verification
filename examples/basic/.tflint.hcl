plugin "terraform" {
  enabled = false
}

plugin "module-signature" {
  enabled = true
}

rule "module_signature_local_source" {
	enabled = true
  allow = true
}

rule "module_signature_registry_source" {
	enabled = true

  denied_module {
    source = "ringanta/backend"
    versions = [
      "< 1.0.0"
    ]
	}

	allowed_module {
    source = "ringanta/backend"
    versions = [
      "~> 1.0"
    ]
	}
}
