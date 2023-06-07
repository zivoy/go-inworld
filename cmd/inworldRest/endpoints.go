package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerEndpoints(e *gin.Engine) {
	e.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	e.GET("/events", toImplement)
	e.GET("/events/:serverId", toImplement)

	s := e.Group("/session")
	{
		s.POST("/open", toImplement)
		s.POST("/:sessionId/message", toImplement)
		s.GET("/:sessionId/custom/:customId", toImplement)

		s.GET("/:sessionId/characters", toImplement)
		s.GET("/:sessionId/character", toImplement)
		s.POST("/:sessionId/character/:characterId", toImplement)

		s.GET("/:sessionId/server/:serverId", toImplement)
		s.GET("/:sessionId/status", toImplement)
		s.GET("/:sessionId/close", toImplement)
		s.GET("/closeall/:uid", toImplement)
		s.GET("/closeall/:uid/server/:serverId", toImplement)
	}

}

func toImplement(c *gin.Context) { //todo
	c.String(http.StatusTeapot, "todo")
}
