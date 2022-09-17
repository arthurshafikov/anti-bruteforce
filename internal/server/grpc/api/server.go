package api

import (
	"context"
	"log"
	"net"

	"github.com/arthurshafikov/anti-bruteforce/internal/server/grpc/generated"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func RunGrpcServer(ctx context.Context, g *errgroup.Group, address string, app App) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}

	a := AppService{
		App: app,
	}
	grpcServer := grpc.NewServer()
	generated.RegisterAppServiceServer(grpcServer, &a)

	g.Go(func() error {
		<-ctx.Done()

		grpcServer.GracefulStop()

		return nil
	})

	g.Go(func() error {
		return grpcServer.Serve(lis)
	})
}
