plugin "terraform" {
  enabled = false
}

plugin "module-signature" {
  enabled = true
}

rule "module_signature_local_source" {
  enabled = true
  allow = false
}
