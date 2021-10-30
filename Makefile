aws:
	cd ./lambda; go get; GOOS=linux CGO_ENABLED=0 go build main.go ; zip lambda.zip main
	mv ./lambda/lambda.zip .
