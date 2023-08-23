package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/shirrashko/BuildingAServer-step2/pkg/api/model"
	"github.com/shirrashko/BuildingAServer-step2/pkg/bl/profile"
	"net/http"
	"strconv"
)

type Handler struct {
	service *profile.Service
}

func NewHandler(s *profile.Service) Handler {
	return Handler{s}
}

func (h Handler) getUserProfileByID(c *gin.Context) {
	// Retrieve the id path parameter from the URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	// Check if the user exists in the database
	if !h.service.IsUserInDB(id) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	// Get the user's profile by ID
	profile, err := h.service.GetProfileByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving profile"})
		return
	}

	// Respond with the profile
	c.IndentedJSON(http.StatusOK, profile) // put the requested profile in the response body
}

// update an existing resource with new data.
func (h Handler) updateUserProfileByID(c *gin.Context) {
	var updatedProfile model.UserProfile
	// check if the given user to add is valid (or in a valid format)
	if err := c.ShouldBindJSON(&updatedProfile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := strconv.Atoi(c.Param("id"))
	err := h.service.UpdateUserProfile(userID, updatedProfile)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found or error updating profile"})
		return
	}

	updatedProfile.ID = userID

	c.JSON(http.StatusOK, updatedProfile)
}

func (h Handler) createUserProfile(c *gin.Context) {
	var newProfile model.UserProfile
	if err := c.BindJSON(&newProfile); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
<<<<<<< HEAD
	h.service.CreateNewProfile(newProfile)
=======

	newID, err := h.service.CreateNewProfile(newProfile)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	newProfile.ID = newID

>>>>>>> temp
	c.IndentedJSON(http.StatusCreated, newProfile)
}
