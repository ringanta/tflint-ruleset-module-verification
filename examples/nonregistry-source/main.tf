module "local_fail" {
  source = "../.."
}

module "local_success" {
  source = "../../terraform-modules"
}

module "github_fail" {
  source = "git@github.com:untrusted/example-module"
}

module "github_success" {
  source = "git@github.com:example/example-module"
}
