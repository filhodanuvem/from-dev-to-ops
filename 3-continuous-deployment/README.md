# 3 - Continuous Deployment with FluxCD and Helm
### Problem

On this chapter you should focus on how to deliver systems into a kubernetes cluster. The technologies I have chosen for that are Flux as a continuous delivery system and Helm as a package manager for k8s. 

If you are not familiar with Docker, this is the best moment to get into it. You should focus on three things: 
- How to create a docker image that contains everything needed to run one small piece of software.
- How to use the image to run a container *. 
- How to upload the image to a docker registry, most common one is [dockerhub](https://hub.docker.com/).


\* A common mistake of beginners on Docker is to misunderstand the usage of volumes. Do not run the container pointing to a volume where the dependencies live. You should be able to send the link of your image in the registry to a friend of you, and they would use it to run a container in their own machine. 

<details>
  <summary>Suggested roadmap</summary>

  - [ ] Create a namespace where you will deploy all the pods. 
  - [ ] Deploy a [nats](https://nats.io/) pod in your namespace using the [helm chart](https://github.com/nats-io/k8s/tree/main/helm/charts/nats). 
  - [ ] Use the [nats controller](https://github.com/nats-io/k8s/tree/main/helm/charts/nack) to create a Stream that listens to the subjects `payment.*`
  - [ ] Create a producer service that continuously publishes messages to nats (for example, it publishes a different message every 3 seconds). Publishes to the subjcect/queue `payment.orders`.
  - [ ] Create a docker image for the producer service. 
  - [ ] Publish a image tag to a docker registry.
  - [ ] Create a local helm chart for the producer service and use a helm release to deploy it into the k8s cluster.
  - [ ] Use [Benthos](https://www.benthos.dev/docs/components/inputs/nats_jetstream) to create a service that consumes from the `payment.orders` queue/subject.
</details>
 

### References

* [Building docker images](https://www.katacoda.com/courses/docker/2)
* [Deploying your first container](https://www.katacoda.com/courses/docker/deploying-first-container)
* [Benthos](https://www.benthos.dev/docs/components/inputs/nats_jetstream)