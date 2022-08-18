package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wenealves10/tracking-app-golang/configs"
	_packageClient "github.com/wenealves10/tracking-app-golang/pkg/client"
)

func init() {
	configs.StartConfigs(".env")
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/", "public/index.html")

	pc, err := _packageClient.NewRabbitClient(configs.URL_RABBITMQ)

	if err != nil {
		panic(err)
	}

	defer pc.Close()

}
