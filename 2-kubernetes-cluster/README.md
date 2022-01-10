# 2 - Kubernetes Cluster


### Problem

Create a kubernetes cluster on AWS EKS. The cluster should contains a flux v2 system running and watching this own repository but on the path `./3-continuous-deployment/flux`. 
Alternatively you can use minikube or kind to set up a cluster locally but I recommend you to do that in AWS to enable terraform run via pipelines (eg.: Terraform Cloud). 

### References 
* [Terraform eks module](https://registry.terraform.io/modules/terraform-aws-modules/eks/aws/latest)
* [flux](https://fluxcd.io/docs/)


