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

type ServerHandler interface {
	Home(*gin.Context)
	Authorize(*gin.Context)
	ResetBucket(*gin.Context)
	AddToWhitelist(*gin.Context)
	AddToBlacklist(*gin.Context)
	RemoveFromWhitelist(*gin.Context)
	RemoveFromBlacklist(*gin.Context)
}

type Server struct {
	httpSrv *http.Server
	Engine  *gin.Engine
	handler ServerHandler
}

func NewServer(handler ServerHandler) *Server {
	return &Server{
		Engine:  gin.Default(),
		handler: handler,
	}
}

func (s *Server) Serve(ctx context.Context, g *errgroup.Group, address string) {
	s.InitRoutes()

	s.httpSrv = &http.Server{
		Addr:    address,
		Handler: s.Engine,
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

func (s *Server) InitRoutes() {
	s.Engine.GET("/", s.handler.Home)
	s.Engine.POST("/authorize", s.handler.Authorize)
	s.Engine.POST("/bucket/reset", s.handler.ResetBucket)
	s.Engine.POST("/whitelist/add", s.handler.AddToWhitelist)
	s.Engine.DELETE("/whitelist/remove", s.handler.RemoveFromWhitelist)
	s.Engine.POST("/blacklist/add", s.handler.AddToBlacklist)
	s.Engine.DELETE("/blacklist/remove", s.handler.RemoveFromBlacklist)
}
