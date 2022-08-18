package usecase

import (
	"context"
	"encoding/json"

	"github.com/wenealves10/tracking-app-golang/domain"
)

type packageUsecase struct {
	pc domain.PackageClient
}

func NewPackageUseCase(pClient domain.PackageClient) domain.PackageUseCase {
	return &packageUsecase{
		pc: pClient,
	}
}

func (p *packageUsecase) TrackByVehicleID(ctx context.Context, id string) (*domain.Package, error) {
	bytes, err := p.pc.ConsumeByVehicleID(ctx, id)

	if err != nil {
		return nil, err
	}

	var res domain.Package

	err = json.Unmarshal(bytes, &res)

	if err != nil {
		return nil, err
	}

	return &res, err

}
