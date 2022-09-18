package api

import (
	"context"
	"log"
	"net"

	"github.com/arthurshafikov/anti-bruteforce/internal/services"
	"github.com/arthurshafikov/anti-bruteforce/internal/transport/grpc/generated"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func RunGrpcServer(ctx context.Context, g *errgroup.Group, address string, services *services.Services) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}

	appService := NewAppService(services)
	grpcServer := grpc.NewServer()
	generated.RegisterAppServiceServer(grpcServer, appService)

	g.Go(func() error {
		<-ctx.Done()

		grpcServer.GracefulStop()

		return nil
	})

	g.Go(func() error {
		return grpcServer.Serve(lis)
	})
}
