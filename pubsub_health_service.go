package pubsub

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/sirupsen/logrus"
)

type PermissionType int

const (
	PermissionPublish   PermissionType = 0
	PermissionSubscribe PermissionType = 1
)

type PubSubHealthService struct {
	name           string
	client         *pubsub.Client
	timeout        time.Duration
	permissionType PermissionType
	resourceId     string
}

func NewPubSubHealthService(name string, client *pubsub.Client, timeout time.Duration, permissionType PermissionType, resourceId string) *PubSubHealthService {
	return &PubSubHealthService{name, client, timeout, permissionType, resourceId}
}

func (h *PubSubHealthService) Name() string {
	return h.name
}

func (h *PubSubHealthService) Check(ctx context.Context) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	var permissions []string
	var err error

	timeoutCtx, _ := context.WithTimeout(ctx, h.timeout)
	if h.permissionType == PermissionPublish {
		permissions, err = h.client.Topic(h.resourceId).IAM().TestPermissions(timeoutCtx, []string{"pubsub.topics.publish"})
	} else if h.permissionType == PermissionSubscribe {
		permissions, err = h.client.Subscription(h.resourceId).IAM().TestPermissions(timeoutCtx, []string{"pubsub.subscriptions.consume"})
	}

	if err != nil {
		logrus.Errorf("Can't TestPermissions %h: %h", h.resourceId, err.Error())
		return res, err
	} else if len(permissions) != 1 {
		return res, fmt.Errorf("invalid permissions: %v", permissions)
	} else {
		return res, nil
	}
}

func (h *PubSubHealthService) Build(ctx context.Context, data map[string]interface{}, err error) map[string]interface{} {
	if err == nil {
		return data
	}
	data["error"] = err.Error()
	return data
}
