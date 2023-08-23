package profile

import (
	"github.com/gin-gonic/gin"
)

func (h Handler) SetUpRoutes(router *gin.Engine) {
	// all endpoints are with the same prefix "/profile/users"
	profilesRouter := router.Group("/profile/users")
<<<<<<< HEAD
	profilesRouter.GET("/:id", h.getAUserProfile)   // get user profile
	profilesRouter.POST("", h.createUserProfile)    // create user profile
	profilesRouter.PUT("/:id", h.updateUserProfile) // update user profile
=======
	profilesRouter.GET("/:id", h.getUserProfileByID)    // get user profile by id
	profilesRouter.POST("", h.createUserProfile)        // create user profile, user data is in the body?
	profilesRouter.PUT("/:id", h.updateUserProfileByID) // update user profile by id
	// todo: need to add delete and patch?
>>>>>>> temp
}
