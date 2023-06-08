package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerEndpoints(e *gin.Engine) {
	e.GET("/", home)
	e.GET("/status", status)

	e.GET("/events", toImplement)
	e.GET("/events/:serverId", toImplement)

	s := e.Group("/session")
	{
		s.POST("/open", openSession)
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

// Deprecated
func toImplement(c *gin.Context) { //todo
	c.String(http.StatusTeapot, "todo")
}

func home(c *gin.Context) {
	c.String(http.StatusOK, "Inworld.AI RESTful Server")
}

func status(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func openSession(c *gin.Context) {
	req := new(OpenSessionRequest)
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := GetApp().OpenSession(req)
	if err != nil {
		c.JSON(http.StatusConflict, "Unable to create session")
		return
	}

	c.JSON(http.StatusOK, res)
}
