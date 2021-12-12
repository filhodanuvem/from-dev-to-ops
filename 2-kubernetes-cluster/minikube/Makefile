minikube-install:
	@curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-arm64
	@sudo install minikube-darwin-arm64 /usr/local/bin/minikube

flux-install:
	wget https://github.com/fluxcd/flux2/releases/download/v0.23.0/flux_0.23.0_darwin_amd64.tar.gz -O flux.tar.gz && \
		tar -xvf flux.tar.gz && \
		sudo mv flux /usr/local/bin/flux && \
		rm flux.tar.gz

minikube-start:
	@minikube start --cpus='4' --memory='4g'

stop:
	@kubectl delete -k ../../3-continuous-deployment/flux
	@flux uninstall --namespace=flux-system
	@minikube stop

start: minikube-start
	@flux install --namespace=flux-system
	flux create source git from-dev-to-ops \
	--branch=master \
	--url=https://github.com/filhodanuvem/from-dev-to-ops \
	--interval=1m
	flux create kustomization from-dev-to-ops \
  --source=from-dev-to-ops \
  --path="./3-continuous-deployment/flux" \
  --interval=10m \
  --health-check-timeout=2m
