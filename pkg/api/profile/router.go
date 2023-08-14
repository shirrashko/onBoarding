package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/shirrashko/BuildingAServer-step2/pkg/api"
)

func (h *Handler) SetUpRoutes(router *gin.Engine) {
	// all endpoints are with the same prefix "/profile/users"
	router.GET("/health", api.HealthCheck) // Associate the GET HTTP method and /health path with a handler function
	//todo: check if it's supposed to be handled by a handler as well
	profilesRouter := router.Group("/profile/users")
	profilesRouter.GET("/:id", h.createUserProfile)  // get user profile
	profilesRouter.POST("", h.updateUserProfile)     // create user profile
	profilesRouter.PUT("/:id", h.getUserProfileByID) // update user profile
}
