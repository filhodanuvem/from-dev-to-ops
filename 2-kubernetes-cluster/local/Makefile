kind-install:
	@go install sigs.k8s.io/kind@v0.11.1

flux-install:
	wget https://github.com/fluxcd/flux2/releases/download/v0.23.0/flux_0.23.0_darwin_amd64.tar.gz -O flux.tar.gz && \
		tar -xvf flux.tar.gz && \
		sudo mv flux /usr/local/bin/flux && \
		rm flux.tar.gz

cilium-install:
	curl -L --remote-name-all https://github.com/cilium/cilium-cli/releases/latest/download/cilium-darwin-amd64.tar.gz{,.sha256sum} && \
		shasum -a 256 -c cilium-darwin-amd64.tar.gz.sha256sum && \
		sudo tar xzvfC cilium-darwin-amd64.tar.gz /usr/local/bin && \
		rm cilium-darwin-amd64.tar.gz{,.sha256sum}

kind-start:
	@kind create cluster --image=kindest/node:v1.23.0 --config=kind.yaml

cilium-start:
	@cilium install
	@cilium status --wait


stop:
	@kind delete cluster

start: kind-start cilium-start
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

kubefwd:
	sudo kubefwd svc -n filhodanuvem-services