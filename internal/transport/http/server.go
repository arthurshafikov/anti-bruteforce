package http

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type Handler interface {
	InitRoutes(engine *gin.Engine)
}

type Server struct {
	httpSrv *http.Server
	handler Handler
}

func NewServer(handler Handler) *Server {
	return &Server{
		handler: handler,
	}
}

func (s *Server) Serve(ctx context.Context, g *errgroup.Group, address string) {
	engine := gin.Default()
	s.handler.InitRoutes(engine)

	s.httpSrv = &http.Server{
		Addr:    address,
		Handler: engine,
	}

	g.Go(func() error {
		<-ctx.Done()

		return s.shutdown()
	})

	if err := s.httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Println("Could not start listener ", err)
	}
}

func (s *Server) shutdown() error {
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.httpSrv.Shutdown(ctx)
}
