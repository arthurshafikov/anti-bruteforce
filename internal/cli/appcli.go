package cli

import (
	"context"
	"time"

	"github.com/arthurshafikov/anti-bruteforce/internal/transport/grpc/api"
	"github.com/arthurshafikov/anti-bruteforce/internal/transport/grpc/generated"
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

func (app *AppCli) AddToWhitelist(subnet string) error {
	_, err := app.grpcClient.AddToWhitelist(app.ctx, &generated.SubnetRequest{
		Subnet: subnet,
	})
	if err != nil {
		app.cancel()
		return err
	}

	return nil
}

func (app *AppCli) RemoveFromWhitelist(subnet string) error {
	_, err := app.grpcClient.RemoveFromWhitelist(app.ctx, &generated.SubnetRequest{
		Subnet: subnet,
	})
	if err != nil {
		app.cancel()
		return err
	}

	return nil
}

func (app *AppCli) AddToBlacklist(subnet string) error {
	_, err := app.grpcClient.AddToBlacklist(app.ctx, &generated.SubnetRequest{
		Subnet: subnet,
	})
	if err != nil {
		app.cancel()
		return err
	}

	return nil
}

func (app *AppCli) RemoveFromBlacklist(subnet string) error {
	_, err := app.grpcClient.RemoveFromBlacklist(app.ctx, &generated.SubnetRequest{
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
