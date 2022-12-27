terraform {
  required_version = "1.2.1"
}

module "registry_fail" {
  source  = "ringanta/backend/aws"
  version = "1.0.0"
}

module "local_fail" {
  source = "../.."
}

module "remote_fail" {
  source = "s3::https://example-terraform-modules.s3.ap-southeast-1.amazonaws.com/vpc.zip"
}
