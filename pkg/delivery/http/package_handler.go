package http

import (
	"context"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/wenealves10/tracking-app-golang/domain"
)

type PackageHandler struct {
	upgrader websocket.Upgrader
	Pusecase domain.PackageUseCase
}

func NewPackageHandler(e *echo.Echo, pu domain.PackageUseCase) {
	handler := &PackageHandler{
		upgrader: websocket.Upgrader{},
		Pusecase: pu,
	}

	e.GET("/package/track/:vehicleID", handler.TrackByVehicleID)
}

func (p *PackageHandler) TrackByVehicleID(c echo.Context) error {
	wsConn, err := p.upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		return err
	}

	ctx, cancelFunc := context.WithCancel(context.Background())

	go func() {
		_, _, err := wsConn.ReadMessage()
		if err != nil {
			cancelFunc()
		}
	}()

	for {
		select {
		case <-ctx.Done():
			wsConn.Close()
			return nil
		default:
			p, err := p.Pusecase.TrackByVehicleID(ctx, c.Param("vehicleID"))
			if err != nil {
				c.Logger().Error(err)
				continue
			}

			err = wsConn.WriteJSON(p)
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}
}
