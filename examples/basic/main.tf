terraform {
  required_version = "1.2.1"
}

module "fail" {
  source  = "ringanta/backend/aws"
  version = "1.0.0"
}

module "success" {
  source = "../.."
}