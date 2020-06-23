package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"os"
)

func NewPubSubClient(ctx context.Context, projectId string, keyFilename string) (*pubsub.Client, error) {
	if len(keyFilename) > 0 && existFile(keyFilename) {
		if logrus.IsLevelEnabled(logrus.InfoLevel) {
			logrus.Info("key file exists")
		}
		return pubsub.NewClient(ctx, projectId, option.WithCredentialsFile(keyFilename))
	} else {
		if logrus.IsLevelEnabled(logrus.WarnLevel) && len(keyFilename) > 0{
			logrus.Warn("key file doesn't exists")
		}
		return pubsub.NewClient(ctx, projectId)
	}
}

func existFile(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		logrus.Error(err)
	}
	return false
}
