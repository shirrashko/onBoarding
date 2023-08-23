package api

import "github.com/gin-gonic/gin"

type handler interface {
	SetUpRoutes(router *gin.Engine)
}
