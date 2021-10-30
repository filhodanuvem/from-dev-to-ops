### AWS localstack study

list of commands to check resources locally 

```bash 
export LOCAL_URL=http://localhost:4566
export AWS_REGION=eu-west-2 

aws --endpoint-url=$LOCAL_URL ec2 describe-vpcs --region $AWS_REGION
```