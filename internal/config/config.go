package config

// App - config app
type App struct {
	GrpcAddr string `env:"GRPC_ADDR,required"`
	MailGun struct {
		Domain     string `env:"DOMAIN,required"`
		PrivateKey string `env:"PRIVATEKEY,required"`
		Name       string `env:"NAME,required"`
	} `envPrefix:"MAILGUN_"`

	Kafka struct {
		Dsn               []string `env:"DSN,required"`
		TopicPersonalized string   `env:"PERSONTOPIC,required"`
		TopicCommon       string   `env:"COMMONTOPIC,required"`
	} `envPrefix:"KAFKA_"`
}
