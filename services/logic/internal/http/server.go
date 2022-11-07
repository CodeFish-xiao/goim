package http

import (
	"github.com/Terry-Mao/goim/services/logic/internal"
	"github.com/Terry-Mao/goim/services/logic/internal/conf"

	"github.com/gin-gonic/gin"
)

// Server is http server.
type Server struct {
	engine *gin.Engine
	logic  *internal.Logic
}

// New new a http server.
func New(c *conf.HTTPServer, l *internal.Logic) *Server {
	engine := gin.New()
	engine.Use(loggerHandler, recoverHandler)
	go func() {
		if err := engine.Run(c.Addr); err != nil {
			panic(err)
		}
	}()
	s := &Server{
		engine: engine,
		logic:  l,
	}
	s.initRouter()
	return s
}

func (s *Server) initRouter() {
	group := s.engine.Group("/goim")
	group.POST("/push/keys", s.pushKeys)
	group.POST("/push/mids", s.pushMids)
	group.POST("/push/room", s.pushRoom)
	group.POST("/push/all", s.pushAll)
	group.GET("/online/top", s.onlineTop)
	group.GET("/online/room", s.onlineRoom)
	group.GET("/online/total", s.onlineTotal)
	group.GET("/nodes/weighted", s.nodesWeighted)
	group.GET("/nodes/instances", s.nodesInstances)
}

// Close close the server.
func (s *Server) Close() {

}
