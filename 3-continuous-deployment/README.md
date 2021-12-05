# 3 - Continuous Deployment

### Dependencies 

This chapter depends of the previous one, where you supposed to create a kubernetes cluster with flux running on it.

### Exercise

- [ ] Create a namespace where you will deploy all the pods. 
- [ ] Deploy a [nats](https://nats.io/) pod running in your namespace using the [helm chart](https://github.com/nats-io/k8s/tree/main/helm/charts/nats). 
- [ ] Use the [nats controller](https://github.com/nats-io/k8s/tree/main/helm/charts/nack) to create a Stream that listens to the subjects `payment.*`
- [ ] Create a producer service that continuously publishes messages to nats (for example, it publishes a different message every 3 seconds). Publishes to the subjcect/queue `payment.orders`.
- [ ] Create a dockerfile that runs the producer service. 
- [ ] Use a docker registry like dockerhub to publish a a new tag for the service.
- [ ] Create a helm chart for the producer service and use a helmrelease to deploy it into the k8s cluster.
- [ ] Use [Benthos](https://www.benthos.dev/docs/components/inputs/nats_jetstream) to create a service that consumes from the `payment.orders` queue/subject.

### References
