package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/common-go/mq"
	"github.com/sirupsen/logrus"
)

type Consumer struct {
	Client       *pubsub.Client
	Subscription *pubsub.Subscription
	AckOnConsume bool
}

func NewConsumer(client *pubsub.Client, subscriptionId string, c SubscriptionConfig, ackOnConsume bool) *Consumer {
	subscription := client.Subscription(subscriptionId)
	return &Consumer{Client: client, Subscription: ConfigureSubscription(subscription, c), AckOnConsume: ackOnConsume}
}

func NewConsumerByConfig(ctx context.Context, c ConsumerConfig, ackOnConsume bool) (*Consumer, error) {
	client, err := NewPubSubClient(ctx, c.Client.ProjectId, c.Client.KeyFilename)
	if err != nil {
		return nil, err
	}
	return NewConsumer(client, c.SubscriptionId, c.SubscriptionConfig, ackOnConsume), nil
}

func ConfigureSubscription(subscription *pubsub.Subscription, c SubscriptionConfig) *pubsub.Subscription {
	if c.MaxOutstandingMessages > 0 {
		subscription.ReceiveSettings.MaxOutstandingMessages = c.MaxOutstandingMessages
	}
	if c.NumGoroutines > 0 {
		subscription.ReceiveSettings.NumGoroutines = c.NumGoroutines
	}
	return subscription
}

func (c *Consumer) Consume(ctx context.Context, caller mq.ConsumerCaller) {
	er1 := c.Subscription.Receive(ctx, func(ctx2 context.Context, m *pubsub.Message) {
		if logrus.IsLevelEnabled(logrus.DebugLevel) {
			logrus.Debugf("Received message: %s", m.Data)
		}
		message := mq.Message{
			Id:         m.ID,
			Data:       m.Data,
			Attributes: m.Attributes,
			Raw:        m,
		}
		if c.AckOnConsume {
			m.Ack()
		}
		caller.Call(ctx2, &message, nil)
	})
	if er1 != nil {
		caller.Call(ctx, nil, er1)
	}
}
