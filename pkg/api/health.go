package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shirrashko/BuildingAServer-step2/pkg/bl"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	if bl.HealthChecking() { // calling the bl
		c.JSON(http.StatusOK, gin.H{"status": "ok"}) // in JSON because requests and responses are being made using it
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"}) // status code 500 - internal server error
	}
}
