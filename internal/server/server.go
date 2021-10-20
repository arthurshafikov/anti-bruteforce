package server

import "github.com/gin-gonic/gin"

type Handler interface {
	Home(*gin.Context)
	Authorize(*gin.Context)
	ResetBucket(*gin.Context)
	AddToWhiteList(*gin.Context)
	RemoveFromWhiteList(*gin.Context)
	AddToBlackList(*gin.Context)
	RemoveFromBlackList(*gin.Context)
}

type Server struct {
	address string
	handler Handler
}

func NewServer(address string, handler Handler) *Server {
	return &Server{
		address: address,
		handler: handler,
	}
}

func (s *Server) Serve() {
	router := gin.Default()

	router.GET("/", s.handler.Home)
	router.POST("/authorize", s.handler.Authorize)
	router.POST("/bucket/reset", s.handler.ResetBucket)
	router.POST("/whitelist/add", s.handler.AddToWhiteList)
	router.POST("/whitelist/remove", s.handler.RemoveFromWhiteList)
	router.POST("/blacklist/add", s.handler.AddToBlackList)
	router.POST("/blacklist/remove", s.handler.RemoveFromBlackList)

	router.Run(s.address)
}
