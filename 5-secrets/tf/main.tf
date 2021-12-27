terraform {
  required_providers {
    vault = {
      source = "hashicorp/vault"
      version = "3.1.1"
    }

    kubernetes = {
      source = "hashicorp/kubernetes"
      version = "2.7.1"
    }
  }
}

provider "vault" {
  address = var.vault_address
  token = var.vault_token
}


provider "kubernetes" {
  config_path = var.kubeconfig
  config_context = var.kubeconfig_context
}
