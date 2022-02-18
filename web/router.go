package web

import (
	"github.com/gin-gonic/gin"
	"github.com/masibw/go-docker-template/web/handler"
)

func NewServer(userHandler *handler.UserHandler) (e *gin.Engine) {
	e = gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	v1 := e.Group("/api/v1")

	users := v1.Group("/users")
	users.GET(":id", userHandler.GetUser)

	return
}
