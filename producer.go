package pubsub

import (
	"context"
	"time"

	"cloud.google.com/go/iam"
	"cloud.google.com/go/pubsub"
	"github.com/sirupsen/logrus"
)

var CheckTopicPermission = CheckPermission

type Producer struct {
	Client *pubsub.Client
	Topic  *pubsub.Topic
}

func NewProducer(ctx context.Context, client *pubsub.Client, topicId string, c TopicConfig) *Producer {
	topic := client.Topic(topicId)
	CheckTopicPermission(ctx, topic.IAM(), "pubsub.topics.publish")
	return &Producer{Client: client, Topic: ConfigureTopic(topic, c)}
}

func NewProducerByConfig(ctx context.Context, c ProducerConfig) (*Producer, error) {
	client, err := NewPubSubClient(ctx, c.Client.ProjectId, c.Client.KeyFilename)
	if err != nil {
		return nil, err
	}
	return NewProducer(ctx, client, c.TopicId, c.Topic), nil
}

func ConfigureTopic(topic *pubsub.Topic, c TopicConfig) *pubsub.Topic {
	if c.CountThreshold > 0 {
		topic.PublishSettings.DelayThreshold = time.Duration(c.CountThreshold) * time.Millisecond
	}
	if c.DelayThreshold > 0 {
		topic.PublishSettings.CountThreshold = c.DelayThreshold
	}
	if c.ByteThreshold > 0 {
		topic.PublishSettings.ByteThreshold = c.ByteThreshold
	}
	if c.NumGoroutines > 0 {
		topic.PublishSettings.NumGoroutines = c.NumGoroutines
	}
	return topic
}

func (c *Producer) Produce(ctx context.Context, data []byte, messageAttributes *map[string]string) (string, error) {
	msg := &pubsub.Message{
		Data: data,
	}

	if messageAttributes != nil {
		msg.Attributes = *messageAttributes
	}

	publishResult := c.Topic.Publish(ctx, msg)
	return publishResult.Get(ctx)
}

func CheckPermission(ctx0 context.Context, iam *iam.Handle, permission string) {
	ctx, _ := context.WithTimeout(ctx0, 30*time.Second)

	if logrus.IsLevelEnabled(logrus.InfoLevel) {
		logrus.Infof("Checking permission: %s", permission)
	}
	if permissions, err := iam.TestPermissions(ctx, []string{permission}); err != nil {
		logrus.Fatalf("Can't check permission %v: %s", permission, err.Error())
	} else if len(permissions) > 0 && permissions[0] == permission {
		logrus.Warnf("Permission %v valid", permission)
	} else {
		logrus.Fatalf("Permission %v invalid", permission)
	}
}
