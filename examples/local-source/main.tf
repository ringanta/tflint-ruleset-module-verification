module "local_fail" {
  source = "../.."
}

module "local_success" {
  source = "../../terraform-modules"
}
