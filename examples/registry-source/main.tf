module "registry_fail" {
  source = "ringanta/vpc/aws"
  version = "1.0.0"
}

module "registry_success" {
  source = "ringanta/backend/aws"
  version = "1.0.0"
}
