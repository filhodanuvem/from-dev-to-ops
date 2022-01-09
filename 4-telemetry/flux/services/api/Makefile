GITHUB_SHA ?= latest

push: build
	docker push cloudson/log-api:${GITHUB_SHA}

build: 
	docker build . -t cloudson/log-api:${GITHUB_SHA}

