package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type EnvCofig struct {
	AppPort        string `env:"APP_PORT,required"`
	AppEnv         string `env:"APP_ENV,required"`
	DBHost         string `env:"DB_HOST,required"`
	DBPort         string `env:"DB_PORT,required"`
	DBUser         string `env:"DB_USER,required"`
	DBPassword     string `env:"DB_PASSWORD,required"`
	DBName         string `env:"DB_NAME,required"`
	DBSSLMode      string `env:"DB_SSL_MODE,required"`
	ConsumerKey    string `env:"MPESA_CONSUMER_KEY,required"`
	ConsumerSecret string `env:"MPESA_CONSUMER_SECRET,required"`
	PassKey        string `env:"MPESA_PASSKEY,required"`
	ShortCode      string `env:"BUSINESS_SHORTCODE,required"`
	MpesaUrl       string `env:"MPESA_URL,required"`
	MPesaTokenUrl  string `env:"MPESA_OAUTH_URL,required"`
	CallbackUrl    string `env:"CALL_BACK_URL,required"`
}

func NewEnvConfig() *EnvCofig {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error loading .env file: %v", err)
	}
	config := &EnvCofig{}

	if err := env.Parse(config); err != nil {
		logrus.Fatalf("Error parsing env file : %v", err)
	}
	logrus.Info("Env Config Loaded successfully")
	return config
}
