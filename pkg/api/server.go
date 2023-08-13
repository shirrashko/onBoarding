package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(routerFactory func() Handlers) Server {
	engine := gin.Default()
	server := Server{engine: engine}
	handlers := routerFactory()
	server.SetUp(handlers) //todo: keep track and make sure I managed everything as I should have
	return server
}

func (server *Server) SetUp(handlers Handlers) {
	for _, h := range handlers.Handlers {
		h.SetUpRoutes(server.engine)
	}
}

func (server *Server) ListenAndServe() { // Attach the engine to a http.Server and start the Server.
	err := server.engine.Run("localhost:8080")
	if err != nil {
		return
	}
}
