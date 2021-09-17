terraform {
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "louisbilliet"
    workspaces {
      prefix = "price-comparator-"
    }
  }
}
