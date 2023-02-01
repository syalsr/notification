package config

// App - config app
type App struct {
	GrpcAddr          string `env:"GRPC_ADDR,required"`
	MailGunDomain     string `env:"MAILGUN_DOMAIN,required"`
	MailGunPrivateKey string `env:"MAILGUN_PRIVATEKEY,required"`
	MailGunName       string `env:"MAILGUN_NAME,required"`
	Kafka             struct {
		Dsn               []string `env:"DSN,required"`
		TopicPersonalized string   `env:"PERSONTOPIC,required"`
		TopicCommon       string   `env:"COMMONTOPIC,required"`
	} `envPrefix:"KAFKA_"`
}
