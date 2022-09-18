package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type Handler interface {
	InitRoutes(engine *gin.Engine)
}

type Logger interface {
	Error(err error)
	Info(msg string)
}

type Server struct {
	logger  Logger
	httpSrv *http.Server
	handler Handler
}

func NewServer(logger Logger, handler Handler) *Server {
	return &Server{
		logger:  logger,
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
		s.logger.Error(fmt.Errorf("could not start listener %w", err))
	}
}

func (s *Server) shutdown() error {
	s.logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.httpSrv.Shutdown(ctx)
}
