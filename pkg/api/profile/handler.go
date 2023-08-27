package profile

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/shirrashko/BuildingAServer-step2/pkg/api/profile/model"
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

func (h Handler) getProfileByID(c *gin.Context) {
	// Retrieve the id path parameter from the URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	// Get the user's userProfile by ID
	userProfile, err := h.service.GetProfileByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) { // sql.ErrNoRows is an error returned when a database query returns no rows.
			// it indicates that a query was executed successfully, but the result set was empty. This error can occur
			// when you're trying to retrieve data from the DB using a query and the data you're looking for
			// doesn't exist in the DB.
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving userProfile"})
		return
	}

	// Respond with the userProfile
	c.JSON(http.StatusOK, userProfile) // put the requested userProfile in the response body
}

// update an existing resource with new data.
func (h Handler) updateProfileByID(c *gin.Context) {
	var updatedProfile model.UserProfile
	if err := c.ShouldBindJSON(&updatedProfile); err != nil { // check if the given user to add is valid (or in a valid format)
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

func (h Handler) createProfile(c *gin.Context) {
	var newProfile model.UserProfile
	if err := c.BindJSON(&newProfile); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	newID, err := h.service.CreateNewProfile(newProfile)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	newProfile.ID = newID

	c.JSON(http.StatusCreated, newProfile)
}
