package pubsub

type ClientConfig struct {
	ProjectId   string `mapstructure:"project_id"`
	KeyFilename string `mapstructure:"key_filename"`
}
