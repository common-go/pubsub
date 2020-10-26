package pubsub

type ProducerConfig struct {
	TopicId string       `mapstructure:"topic_id"`
	Client  ClientConfig `mapstructure:"client"`
	Topic   TopicConfig  `mapstructure:"topic"`
	Retry   RetryConfig  `mapstructure:"retry"`
}

type TopicConfig struct {
	DelayThreshold int `mapstructure:"delay_threshold"` // MaxMessages
	CountThreshold int `mapstructure:"count_threshold"` // MaxMilliseconds
	ByteThreshold  int `mapstructure:"byte_threshold"`  // MaxBytes
	NumGoroutines  int `mapstructure:"num_goroutines"`
}
