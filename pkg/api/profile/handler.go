package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/shirrashko/BuildingAServer-step2/pkg/bl"
	"github.com/shirrashko/BuildingAServer-step2/pkg/db"
	"net/http"
	"strconv"
)

type Handler struct {
	service *bl.Service
}

func NewHandler(s *bl.Service) Handler {
	return Handler{s}
}

// implementation of the methods of the Service object, which regard to the db contains users profile info

func (h *Handler) getAUserProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id")) // retrieve the id path parameter from the URL
	if h.service.IsUserInDB(id) {
		c.IndentedJSON(http.StatusOK, h.service.GetProfileByID(id))
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ID not found"})
	}
}

// update an existing resource with new data.
func (h *Handler) updateUserProfile(c *gin.Context) {
	var newUser db.UserProfile
	// check if the given user to add is valid (or in a valid format)
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := strconv.Atoi(c.Param("id"))
	if !h.service.IsUserInDB(userID) {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else { // adds a user from JSON received in the request body
		h.service.UpdateUserProfile(userID, newUser)
		c.JSON(http.StatusOK, newUser)
	}
}

func (h *Handler) createUserProfile(c *gin.Context) {
	var newProfile db.UserProfile
	if err := c.BindJSON(&newProfile); err != nil { // bind the received JSON to newProfile.
		return
	}
	h.service.CreateNewUser(newProfile)
	c.IndentedJSON(http.StatusCreated, newProfile)
}

// parsing the JSON data and mapping it to a Go struct. If the JSON data doesn't match the Go struct or if there
// are validation errors (e.g., wrong data types), Gin provides error messages that you can handle.
