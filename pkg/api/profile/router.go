package profile

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetUpRoutes(router *gin.Engine) {
	// all endpoints are with the same prefix "/profile/users"
	// router.GET("/health", HealthCheck) // Associate the GET HTTP method and /health path with a handler function
	//todo: check if the health check supposed to be handled by a handler as well
	profilesRouter := router.Group("/profile/users")
	profilesRouter.GET("/:id", h.getUserProfileByID)    // get user profile by id
	profilesRouter.POST("", h.createUserProfile)        // create user profile, user data is in the body?
	profilesRouter.PUT("/:id", h.updateUserProfileByID) // update user profile by id
	// todo: need to add delete and patch?
}
