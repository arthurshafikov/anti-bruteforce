package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Home(*gin.Context)
	Authorize(*gin.Context)
	ResetBucket(*gin.Context)
	AddToWhiteList(*gin.Context)
	AddToBlackList(*gin.Context)
	RemoveFromWhiteList(*gin.Context)
	RemoveFromBlackList(*gin.Context)
}

type Server struct {
	httpSrv *http.Server
	address string
	engine  *gin.Engine
	handler Handler
}

func NewServer(address string, handler Handler) *Server {
	return &Server{
		address: address,
		engine:  gin.Default(),
		handler: handler,
	}
}

func (s *Server) Serve(ctx context.Context) {
	s.initRoutes()

	s.httpSrv = &http.Server{
		Addr:    s.address,
		Handler: s.engine,
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

func (s *Server) initRoutes() {
	s.engine.GET("/", s.handler.Home)
	s.engine.POST("/authorize", s.handler.Authorize)
	s.engine.POST("/bucket/reset", s.handler.ResetBucket)
	s.engine.POST("/whitelist/add", s.handler.AddToWhiteList)
	s.engine.DELETE("/whitelist/remove", s.handler.RemoveFromWhiteList)
	s.engine.POST("/blacklist/add", s.handler.AddToBlackList)
	s.engine.DELETE("/blacklist/remove", s.handler.RemoveFromBlackList)
}
