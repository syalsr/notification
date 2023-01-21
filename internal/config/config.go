package config

// App - config app
type App struct {
	GrpcAddr               string `env:"GRPC_ADDR,required"`
	MailGunDomain          string `env:"MAILGUN_DOMAIN,required"`
	MailGunPrivateKey      string `env:"MAILGUN_PRIVATEKEY,required"`
	MailGunName            string `env:"MAILGUN_NAME,required"`
	KafkaURL               []string
	KafkaParition          int32
	KafkaOffset            int64
	KafkaTopicPersonalized string
	KafkaTopicCommon       string
}
