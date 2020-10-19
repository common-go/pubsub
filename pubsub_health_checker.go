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

type PubSubHealthChecker struct {
	name           string
	client         *pubsub.Client
	timeout        time.Duration
	permissionType PermissionType
	resourceId     string
}

func NewPubSubHealthChecker(name string, client *pubsub.Client, timeout time.Duration, permissionType PermissionType, resourceId string) *PubSubHealthChecker {
	return &PubSubHealthChecker{name, client, timeout, permissionType, resourceId}
}

func (h *PubSubHealthChecker) Name() string {
	return h.name
}

func (h *PubSubHealthChecker) Check(ctx context.Context) (map[string]interface{}, error) {
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
		logrus.Errorf("Can't TestPermissions %s: %s", h.resourceId, err.Error())
		return res, err
	} else if len(permissions) != 1 {
		return res, fmt.Errorf("invalid permissions: %v", permissions)
	} else {
		return res, nil
	}
}

func (h *PubSubHealthChecker) Build(ctx context.Context, data map[string]interface{}, err error) map[string]interface{} {
	if err == nil {
		return data
	}
	data["error"] = err.Error()
	return data
}