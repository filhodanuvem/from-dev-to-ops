terraform {
  backend "remote" {
    organization = "from-dev-to-ops-org"
    workspaces {
      name = "from-dev-to-ops-2-kubernetes-cluster"
    }
  }
}