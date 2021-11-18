terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "3.65.0"
    }

    flux = {
      source = "fluxcd/flux"
      version = "0.7.3"
    }

    kubectl = {
      source  = "gavinbunney/kubectl"
      version = ">= 1.10.0"
    }
  }
}

provider "aws" {
  # Configuration options
  region = "eu-west-2"
}


provider "flux" {
  # Configuration options
}

provider "kubectl" {

}


provider "kubernetes" {
  host                   = data.aws_eks_cluster.eks.endpoint
  cluster_ca_certificate = base64decode(data.aws_eks_cluster.eks.certificate_authority[0].data)
  token                  = data.aws_eks_cluster_auth.eks.token
}