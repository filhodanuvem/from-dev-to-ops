
locals {
  cluster_name = "chapter2-eks-${random_string.suffix.result}"
}

resource "random_string" "suffix" {
  length  = 8
  special = false
}

data "aws_eks_cluster" "eks" {
  name = module.eks.cluster_id
}

data "aws_eks_cluster_auth" "eks" {
  name = module.eks.cluster_id
}

module "eks" {
  source          = "terraform-aws-modules/eks/aws"

  cluster_version = "1.21"
  cluster_name    = local.cluster_name
  vpc_id          = module.vpc.vpc_id
  subnets         = module.vpc.public_subnets


  worker_groups = [
     {
      name                          = "worker-group-1"
      instance_type                 = "t2.small"
      spot_price                    = "0.0078"
      additional_userdata           = "echo foo bar"
      asg_desired_capacity          = 2
      additional_security_group_ids = [aws_security_group.worker_group_mgmt_one.id]
    }
  ]
}