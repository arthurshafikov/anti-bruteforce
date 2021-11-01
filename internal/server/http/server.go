package http

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
	address string
	Engine  *gin.Engine
	handler ServerHandler
}

func NewServer(address string, handler ServerHandler) *Server {
	return &Server{
		address: address,
		Engine:  gin.Default(),
		handler: handler,
	}
}

func (s *Server) Serve(ctx context.Context) {
	s.InitRoutes()

	s.httpSrv = &http.Server{
		Addr:    s.address,
		Handler: s.Engine,
	}

	go s.shutdownOnContextDone(ctx)

	if err := s.httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Println("Could not start listener ", err)
	}
}

func (s *Server) shutdownOnContextDone(ctx context.Context) {
	<-ctx.Done()

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := s.httpSrv.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown: ", err)
	}
	cancel()
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
