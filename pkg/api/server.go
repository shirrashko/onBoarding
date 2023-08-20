package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(routerFactory func() (Handlers, error)) (Server, error) {
	engine := gin.Default()
	server := Server{engine: engine}
	handlers, err := routerFactory()
	if err != nil {
		return server, err
	}
	server.SetUp(handlers)
	return server, nil
}

func (server *Server) SetUp(handlers Handlers) {
	for _, h := range handlers.handlers {
		h.SetUpRoutes(server.engine)
	}
}

func (server *Server) ListenAndServe() error { // Attach the engine to a http.Server and start the Server.
	return server.engine.Run("localhost:8080")
}
