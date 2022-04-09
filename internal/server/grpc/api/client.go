package api

import (
	"context"
	"log"

	"github.com/arthurshafikov/anti-bruteforce/internal/server/grpc/generated"
	"google.golang.org/grpc"
)

func NewGRPCClient(ctx context.Context, address string) generated.AppServiceClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		<-ctx.Done()
		conn.Close()
	}()

	return generated.NewAppServiceClient(conn)
}
