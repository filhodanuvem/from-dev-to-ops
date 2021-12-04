# 1 - AWS Networking - Basics with Lambda 

### Dependencies 

This chapter doesn't depend on anything and it's not required to move on to Chapter 2. You may run `terraform destroy` when you complete it.
### Exercise

Create and deploy a lambda function that is not accessible outside of a private network but which has access to https://checkip.amazonaws.com/. Use terraform to manage the resources. 

### Solution

In this chapter I've configured a VPC with 2 subnets, a private and a public one. 
The public subnet is connected to a Internet gateway, giving it access to internet. 
The private one is connected to a NATS gateway, both via route table.