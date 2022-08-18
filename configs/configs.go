package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	USER_RABBITMQ     = ""
	PASSWORD_RABBITMQ = ""
	HOST_RABBITMQ     = ""
	PORT_RABBITMQ     = ""
	URL_RABBITMQ      = ""
)

func StartConfigs(envPath string) {
	err := godotenv.Load(envPath)

	if err != nil {
		panic(err)
	}

	USER_RABBITMQ = os.Getenv("USER_RABBITMQ")
	PASSWORD_RABBITMQ = os.Getenv("PASSWORD_RABBITMQ")
	HOST_RABBITMQ = os.Getenv("HOST_RABBITMQ")
	PORT_RABBITMQ = os.Getenv("PORT_RABBITMQ")

	URL_RABBITMQ = fmt.Sprintf("amqp://%s:%s@%s:%s/", USER_RABBITMQ, PASSWORD_RABBITMQ, HOST_RABBITMQ, PORT_RABBITMQ)
}
