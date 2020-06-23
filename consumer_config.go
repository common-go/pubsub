package pubsub

type ConsumerConfig struct {
	SubscriptionId     string             `mapstructure:"subscription_id"`
	Client             ClientConfig       `mapstructure:"client"`
	SubscriptionConfig SubscriptionConfig `mapstructure:"subscription"`
}

type SubscriptionConfig struct {
	MaxOutstandingMessages int `mapstructure:"max_outstanding_messages"`
	NumGoroutines          int `mapstructure:"num_goroutines"`
}
