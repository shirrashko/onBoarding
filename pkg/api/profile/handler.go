package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/shirrashko/BuildingAServer-step2/pkg/api/model"
	"github.com/shirrashko/BuildingAServer-step2/pkg/bl"
	"net/http"
	"strconv"
)

type Handler struct {
	service *bl.Service
}

func NewHandler(s *bl.Service) Handler {
	return Handler{s}
}

func (h *Handler) getUserProfileByID(c *gin.Context) {
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
func (h *Handler) updateUserProfileByID(c *gin.Context) {
	var newUser model.UserProfile
	// check if the given user to add is valid (or in a valid format)
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := strconv.Atoi(c.Param("id"))
	err := h.service.UpdateUserProfile(userID, newUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found or error updating profile"})
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func (h *Handler) createUserProfile(c *gin.Context) {
	var newProfile model.UserProfile
	if err := c.BindJSON(&newProfile); err != nil { // bind the received JSON to newProfile.
		return
	}
	_ = h.service.CreateNewProfile(newProfile) //todo: handle error
	c.IndentedJSON(http.StatusCreated, newProfile)
}
