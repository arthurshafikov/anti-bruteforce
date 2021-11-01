package api

import (
	"context"

	"github.com/thewolf27/anti-bruteforce/internal/models"
	"github.com/thewolf27/anti-bruteforce/internal/server/grpc/generated"
)

type App interface {
	ResetBucket()
	AddToWhitelist(models.SubnetInput) error
	AddToBlacklist(models.SubnetInput) error
	RemoveFromWhitelist(models.SubnetInput) error
	RemoveFromBlacklist(models.SubnetInput) error
}

var successResponse = &generated.ServerResponse{
	Data: "OK",
}

type AppService struct {
	App App
	generated.UnimplementedAppServiceServer
}

func (a *AppService) ResetBucket(ctx context.Context, req *generated.EmptyRequest) (*generated.ServerResponse, error) {
	a.App.ResetBucket()

	return successResponse, nil
}

func (a *AppService) AddToWhitelist(
	ctx context.Context,
	req *generated.SubnetRequest,
) (*generated.ServerResponse, error) {
	err := a.App.AddToWhitelist(models.SubnetInput{
		Subnet: req.Subnet,
	})
	if err != nil {
		return nil, err
	}

	return successResponse, nil
}

func (a *AppService) AddToBlacklist(
	ctx context.Context,
	req *generated.SubnetRequest,
) (*generated.ServerResponse, error) {
	err := a.App.AddToBlacklist(models.SubnetInput{
		Subnet: req.Subnet,
	})
	if err != nil {
		return nil, err
	}

	return successResponse, nil
}

func (a *AppService) RemoveFromWhitelist(
	ctx context.Context,
	req *generated.SubnetRequest,
) (*generated.ServerResponse, error) {
	err := a.App.RemoveFromWhitelist(models.SubnetInput{
		Subnet: req.Subnet,
	})
	if err != nil {
		return nil, err
	}

	return successResponse, nil
}

func (a *AppService) RemoveFromBlacklist(
	ctx context.Context,
	req *generated.SubnetRequest,
) (*generated.ServerResponse, error) {
	err := a.App.RemoveFromBlacklist(models.SubnetInput{
		Subnet: req.Subnet,
	})
	if err != nil {
		return nil, err
	}

	return successResponse, nil
}
