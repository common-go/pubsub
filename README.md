# pubsub

A fully managed messaging service that allows for event-driven systems and real-time analytics on Google Cloud Platform. Key features include:
- <b>Scalability</b>: Automatically scales to handle high-throughput workloads.
- <b>Durability</b>: Ensures message delivery with at-least-once delivery guarantees.
- <b>Flexibility</b>: Supports both push and pull delivery models.
- <b>Integration</b>: Easily integrates with other Google Cloud services.

### Use Cases of Google Pub/Sub
Common use cases include event-driven architectures, log collection, and streaming analytics.

![Microservice Architecture](https://cdn-images-1.medium.com/max/800/1*vKeePO_UC73i7tfymSmYNA.png)

#### Event-Driven Architectures
- <b>Scenario</b>: Building applications where different components communicate via events (e.g., microservices)
- <b>Benefit</b>: Decouples components, allowing independent scaling and development
![A typical micro service](https://cdn-images-1.medium.com/max/800/1*d9kyekAbQYBxH-C6w38XZQ.png)
#### Log Collection and Monitoring
- <b>Scenario</b>: Aggregating logs from multiple applications and systems.
- <b>Benefit</b>: Centralized logging and monitoring, improving visibility and debugging capabilities.


#### Streaming analytics
- <b>Scenario</b>: Collecting and analyzing data streams from various sources like IoT devices, social media, or user activity.
- <b>Benefit</b>: Enables real-time data processing and analytics, providing timely insights and actions.

### Libraries for Google Pub/Sub
- GO: [pubsub](https://github.com/core-go/pubsub). Example is at [go-pubsub-sample](https://github.com/project-samples/go-pubsub-sample)
- nodejs: [pubsub](https://github.com/core-ts/pubsub). Example is at [pubsub-sample](https://github.com/typescript-tutorial/pubsub-sample)

#### A common flow to consume a message from a message queue
![A common flow to consume a message from a message queue](https://cdn-images-1.medium.com/max/800/1*Y4QUN6QnfmJgaKigcNHbQA.png)
- The libraries to implement this flow are:
  - [mq](https://github.com/core-go/mq) for GOLANG. Example is at [go-subscription](https://github.com/project-samples/go-subscription)
  - [mq-one](https://www.npmjs.com/package/mq-one) for nodejs. Example is at [pubsub-sample](https://github.com/typescript-tutorial/pubsub-sample)

### Comparison of Google Pub/Sub, Amazon SQS, and Apache Kafka
#### Google Pub/Sub:
- <b>Type</b>: Managed real-time messaging service.
- <b>Use Case</b>: Event-driven architectures, real-time analytics.
- <b>Scalability</b>: Automatically scales.
- <b>Delivery Guarantees</b>: At-least-once delivery.
- <b>Integration</b>: Tight with Google Cloud services.
- <b>Delivery Models</b>: Push and pull.

#### Amazon SQS
- <b>Type</b>: Managed message queuing service.
- <b>Use Case</b>: Decoupling and scaling microservices, asynchronous tasks.
- <b>Scalability</b>: Automatically scales.
- <b>Delivery Guarantees</b>: At-least-once, FIFO (exactly-once).
- <b>Integration</b>: Deep integration with AWS services.
- <b>Delivery Models</b>: Primarily pull, with long polling.

#### Apache Kafka
- <b>Type</b>: Open-source event streaming platform.
- <b>Use Case</b>: High-throughput messaging, event sourcing, log aggregation.
- <b>Scalability</b>: High with partitioned topics.
- <b>Delivery Guarantees</b>: Configurable (at-least-once, exactly-once).
- <b>Integration</b>: Broad ecosystem with various connectors.
- <b>Delivery Models</b>: Pull-based consumer groups.

### Key Differences
- <b>Management</b>: Pub/Sub and SQS are managed services, while Kafka is typically self-managed or via managed services like Confluent.
- <b>Use Case Focus</b>: Pub/Sub and Kafka are ideal for real-time processing, whereas SQS is great for decoupling microservices and handling asynchronous tasks.
- <b>Delivery Models</b>: Pub/Sub supports push and pull, SQS supports pull with long polling, and Kafka primarily uses pull with consumer groups.
- <b>Scalability</b>: All three are highly scalable, but Kafka offers the most control over performance tuning.
- <b>Integration</b>: Pub/Sub integrates well with Google Cloud, SQS with AWS, and Kafka has a broad integration ecosystem.

### When to Use
- <b>Google Pub/Sub</b>: If you're using Google Cloud and need a managed, real-time messaging solution.
- <b>Amazon SQS</b>: For reliable, scalable message queuing in AWS environments.
- <b>Apache Kafka</b>: For complex event streaming and log aggregation, with a need for fine-tuned control and a broad integration ecosystem.

## Installation

Please make sure to initialize a Go module before installing core-go/pubsub:

```shell
go get -u github.com/core-go/pubsub
```

Import:

```go
import "github.com/core-go/pubsub"
```
