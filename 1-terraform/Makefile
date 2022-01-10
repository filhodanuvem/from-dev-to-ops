all: aws 

pre-aws:
	rm lambda.zip
	rm ./lambda/main

aws: pre-aws
	cd ./lambda; go get; GOOS=linux CGO_ENABLED=0 go build main.go ; zip lambda.zip main
	mv ./lambda/lambda.zip .
