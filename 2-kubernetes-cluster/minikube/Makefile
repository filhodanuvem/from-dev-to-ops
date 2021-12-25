kind-install:
	@go install sigs.k8s.io/kind@v0.11.1

flux-install:
	wget https://github.com/fluxcd/flux2/releases/download/v0.23.0/flux_0.23.0_darwin_amd64.tar.gz -O flux.tar.gz && \
		tar -xvf flux.tar.gz && \
		sudo mv flux /usr/local/bin/flux && \
		rm flux.tar.gz

kind-start:
	@kind create cluster --image=kindest/node:v1.23.0

stop:
	@kubectl delete -k ../../3-continuous-deployment/flux
	@flux uninstall --namespace=flux-system
	@kind delete cluster

start: kind-start
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
