package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wenealves10/tracking-app-golang/configs"
	"github.com/wenealves10/tracking-app-golang/domain"
	_packageClient "github.com/wenealves10/tracking-app-golang/pkg/client"
	_packageHandler "github.com/wenealves10/tracking-app-golang/pkg/delivery/http"
	_packageUsecase "github.com/wenealves10/tracking-app-golang/pkg/usecase"
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

	dataJson, err := os.ReadFile("./data/generated.json")
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	var packages []domain.Package

	err = json.Unmarshal(dataJson, &packages)

	if err != nil {
		panic(err)
	}

	go func() {
		for _, pkg := range packages {
			time.Sleep(10 * time.Microsecond)
			pc.Publish(domain.Package{From: pkg.From, To: pkg.To, VehicleID: "123"})
		}
	}()

	pu := _packageUsecase.NewPackageUseCase(pc)

	_packageHandler.NewPackageHandler(e, pu)

	e.Logger.Fatal(e.Start(":3000"))
}
