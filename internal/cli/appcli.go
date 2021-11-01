package cli

import (
	"context"
	"time"

	"github.com/thewolf27/anti-bruteforce/internal/server/grpc/api"
	"github.com/thewolf27/anti-bruteforce/internal/server/grpc/generated"
)

type AppCli struct {
	ctx        context.Context
	cancel     context.CancelFunc
	grpcClient generated.AppServiceClient
}

func NewAppCli(address string) *AppCli {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	return &AppCli{
		ctx:        ctx,
		cancel:     cancel,
		grpcClient: api.NewGRPCClient(ctx, address),
	}
}

func (app *AppCli) AddToWhiteList(subnet string) error {
	_, err := app.grpcClient.AddToWhiteList(app.ctx, &generated.SubnetRequest{
		Subnet: subnet,
	})
	if err != nil {
		app.cancel()
		return err
	}

	return nil
}

func (app *AppCli) RemoveFromWhiteList(subnet string) error {
	_, err := app.grpcClient.RemoveFromWhiteList(app.ctx, &generated.SubnetRequest{
		Subnet: subnet,
	})
	if err != nil {
		app.cancel()
		return err
	}

	return nil
}

func (app *AppCli) AddToBlackList(subnet string) error {
	_, err := app.grpcClient.AddToBlackList(app.ctx, &generated.SubnetRequest{
		Subnet: subnet,
	})
	if err != nil {
		app.cancel()
		return err
	}

	return nil
}

func (app *AppCli) RemoveFromBlackList(subnet string) error {
	_, err := app.grpcClient.RemoveFromBlackList(app.ctx, &generated.SubnetRequest{
		Subnet: subnet,
	})
	if err != nil {
		app.cancel()
		return err
	}

	return nil
}

func (app *AppCli) ResetBucket() error {
	_, err := app.grpcClient.ResetBucket(app.ctx, &generated.EmptyRequest{})
	if err != nil {
		app.cancel()
		return err
	}

	return nil
}
