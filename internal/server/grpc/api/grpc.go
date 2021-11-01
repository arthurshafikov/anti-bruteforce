package api

import (
	"context"

	"github.com/thewolf27/anti-bruteforce/internal/models"
	"github.com/thewolf27/anti-bruteforce/internal/server/grpc/generated"
)

type App interface {
	ResetBucket()
	AddToWhiteList(models.SubnetInput) error
	AddToBlackList(models.SubnetInput) error
	RemoveFromWhiteList(models.SubnetInput) error
	RemoveFromBlackList(models.SubnetInput) error
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

func (a *AppService) AddToWhiteList(
	ctx context.Context,
	req *generated.SubnetRequest,
) (*generated.ServerResponse, error) {
	err := a.App.AddToWhiteList(models.SubnetInput{
		Subnet: req.Subnet,
	})
	if err != nil {
		return nil, err
	}

	return successResponse, nil
}

func (a *AppService) AddToBlackList(
	ctx context.Context,
	req *generated.SubnetRequest,
) (*generated.ServerResponse, error) {
	err := a.App.AddToBlackList(models.SubnetInput{
		Subnet: req.Subnet,
	})
	if err != nil {
		return nil, err
	}

	return successResponse, nil
}

func (a *AppService) RemoveFromWhiteList(
	ctx context.Context,
	req *generated.SubnetRequest,
) (*generated.ServerResponse, error) {
	err := a.App.RemoveFromWhiteList(models.SubnetInput{
		Subnet: req.Subnet,
	})
	if err != nil {
		return nil, err
	}

	return successResponse, nil
}

func (a *AppService) RemoveFromBlackList(
	ctx context.Context,
	req *generated.SubnetRequest,
) (*generated.ServerResponse, error) {
	err := a.App.RemoveFromBlackList(models.SubnetInput{
		Subnet: req.Subnet,
	})
	if err != nil {
		return nil, err
	}

	return successResponse, nil
}
