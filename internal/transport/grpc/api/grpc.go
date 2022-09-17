package api

import (
	"context"

	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	"github.com/arthurshafikov/anti-bruteforce/internal/transport/grpc/generated"
)

type App interface {
	ResetBucket()
	AddToWhitelist(core.SubnetInput) error
	AddToBlacklist(core.SubnetInput) error
	RemoveFromWhitelist(core.SubnetInput) error
	RemoveFromBlacklist(core.SubnetInput) error
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
	err := a.App.AddToWhitelist(core.SubnetInput{
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
	err := a.App.AddToBlacklist(core.SubnetInput{
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
	err := a.App.RemoveFromWhitelist(core.SubnetInput{
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
	err := a.App.RemoveFromBlacklist(core.SubnetInput{
		Subnet: req.Subnet,
	})
	if err != nil {
		return nil, err
	}

	return successResponse, nil
}
