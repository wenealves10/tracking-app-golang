package domain

import "context"

type Package struct {
	From      string `json:"from"`
	To        string `json:"to"`
	VehicleID string `json:"vehicleID"`
}

type PackageUseCase interface {
	TrackByVehicleID(ctx context.Context, id string) (*Package, error)
}

type PackageClient interface {
	ConsumeByVehicleID(ctx context.Context, vehicleID string) ([]byte, error)
}
