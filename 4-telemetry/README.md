# 4 - Telemetry

### Problem

This section is focused in observability using different data such as metrics, logs and tracing. 
Use tracing to connect *events* in different services.

#### Tracing 

```
Change the producer and consumer deployed to use generate a chain of tracing spans. You might need to deploy one more service, it can be an api that receive requests from the consumer.   
```

#### Metrics 

```
Change one of the services to emit a custom metric that will be pulled by a prometheus server.
After that deploy a grafana cluster and create a dashboard that uses the metric from prometheus.
```

### References 

* [OpenTelemetry](https://opentelemetry.io/)
* [Prometheus scraping pods](https://www.weave.works/docs/cloud/latest/tasks/monitor/configuration-k8s/)