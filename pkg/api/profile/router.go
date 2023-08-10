package profile

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetUpRoutes(router *gin.Engine) {
	// all endpoints are with the same prefix "/profile/users"
	profilesRouter := router.Group("/profile/users")
	profilesRouter.GET("/:id", h.createUserProfile) // get user profile
	profilesRouter.POST("", h.updateUserProfile)    // create user profile
	profilesRouter.PUT("/:id", h.getAUserProfile)   // update user profile
}
