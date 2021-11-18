
data "aws_eks_cluster" "eks" {
  name = module.eks.cluster_id
}

data "aws_eks_cluster_auth" "eks" {
  name = module.eks.cluster_id
}

module "eks" {
  source          = "terraform-aws-modules/eks/aws"

  cluster_version = "1.21"
  cluster_name    = "chapeter-2-cluster"
  vpc_id          = aws_vpc.main.id 
  subnets         = [aws_subnet.main_k8s_subnet_1.id, aws_subnet.main_k8s_subnet_2.id]

  worker_groups = [
    {
      instance_type = "m4.large"
      asg_max_size  = 5
      asg_desired_capacity = 2
    }
  ]
}