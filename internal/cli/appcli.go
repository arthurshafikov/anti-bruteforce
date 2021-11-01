package cli

import (
	"context"

	"github.com/thewolf27/anti-bruteforce/internal/server/grpc/api"
	"github.com/thewolf27/anti-bruteforce/internal/server/grpc/generated"
)

type AppCli struct {
	ctx        context.Context
	grpcClient generated.AppServiceClient
}

func NewAppCli(address string) *AppCli {
	ctx := context.Background()

	return &AppCli{
		ctx:        ctx,
		grpcClient: api.NewGRPCClient(ctx, address),
	}
}

func (app *AppCli) AddToWhiteList(subnet string) error {
	_, err := app.grpcClient.AddToWhiteList(app.ctx, &generated.SubnetRequest{
		Subnet: subnet,
	})
	if err != nil {
		return err
	}

	return nil
}

func (app *AppCli) RemoveFromWhiteList(subnet string) error {
	_, err := app.grpcClient.RemoveFromWhiteList(app.ctx, &generated.SubnetRequest{
		Subnet: subnet,
	})
	if err != nil {
		return err
	}

	return nil
}

func (app *AppCli) AddToBlackList(subnet string) error {
	_, err := app.grpcClient.AddToBlackList(app.ctx, &generated.SubnetRequest{
		Subnet: subnet,
	})
	if err != nil {
		return err
	}

	return nil
}

func (app *AppCli) RemoveFromBlackList(subnet string) error {
	_, err := app.grpcClient.RemoveFromBlackList(app.ctx, &generated.SubnetRequest{
		Subnet: subnet,
	})
	if err != nil {
		return err
	}

	return nil
}

func (app *AppCli) ResetBucket() error {
	_, err := app.grpcClient.ResetBucket(app.ctx, &generated.EmptyRequest{})
	if err != nil {
		return err
	}

	return nil
}
