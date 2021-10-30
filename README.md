### AWS localstack study

list of commands to check resources locally 

```bash 
export LOCAL_URL=http://localhost:4566
export AWS_REGION=eu-west-2 

aws --endpoint-url=$LOCAL_URL ec2 describe-vpcs --region $AWS_REGION
aws lambda invoke --function-name lambda_function_name --payload '{"name":"Bob"}' --invocation-type Event --cli-binary-format raw-in-base64-out  ./output.json --endpoint-url=$LOCAL_URL
```