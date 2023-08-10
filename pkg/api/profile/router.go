package profile

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetUpRoutes(router *gin.Engine) {
	// all endpoints are with the same prefix "/profile/users"
	profilesRouter := router.Group("/profile/users")
	profilesRouter.GET("/:id", h.getAUserProfile)   // get user profile
	profilesRouter.POST("", h.createUserProfile)    // create user profile
	profilesRouter.PUT("/:id", h.updateUserProfile) // update user profile
}
