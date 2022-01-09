GITHUB_SHA ?= latest

push: build
	docker push cloudson/producer:${GITHUB_SHA}

build: 
	docker build . -t cloudson/producer:${GITHUB_SHA}

