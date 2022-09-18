package api

import (
	"context"

	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	"github.com/arthurshafikov/anti-bruteforce/internal/services"
	"github.com/arthurshafikov/anti-bruteforce/internal/transport/grpc/generated"
)

var successResponse = &generated.ServerResponse{
	Data: "OK",
}

type AppService struct {
	services *services.Services
	generated.UnimplementedAppServiceServer
}

func NewAppService(services *services.Services) *AppService {
	return &AppService{
		services: services,
	}
}

func (a *AppService) ResetBucket(ctx context.Context, req *generated.EmptyRequest) (*generated.ServerResponse, error) {
	a.services.Bucket.ResetBucket()

	return successResponse, nil
}

func (a *AppService) AddToWhitelist(
	ctx context.Context,
	req *generated.SubnetRequest,
) (*generated.ServerResponse, error) {
	err := a.services.Whitelist.AddToWhitelist(core.SubnetInput{
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
	err := a.services.Blacklist.AddToBlacklist(core.SubnetInput{
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
	err := a.services.Whitelist.RemoveFromWhitelist(core.SubnetInput{
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
	err := a.services.Blacklist.RemoveFromBlacklist(core.SubnetInput{
		Subnet: req.Subnet,
	})
	if err != nil {
		return nil, err
	}

	return successResponse, nil
}
