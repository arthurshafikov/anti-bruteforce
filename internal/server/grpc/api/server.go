package api

import (
	"context"
	"log"
	"net"

	"github.com/arthurshafikov/anti-bruteforce/internal/server/grpc/generated"
	"google.golang.org/grpc"
)

func RunGrpcServer(ctx context.Context, address string, app App) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}

	a := AppService{
		App: app,
	}
	grpcServer := grpc.NewServer()
	generated.RegisterAppServiceServer(grpcServer, &a)

	go func() {
		<-ctx.Done()
		grpcServer.Stop()
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
