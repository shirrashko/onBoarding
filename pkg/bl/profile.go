package bl

import (
	"github.com/shirrashko/BuildingAServer-step2/pkg/db"
)

// Service The service has a repository with a client field, which is the connection to the database we are working with
type Service struct {
	repository *db.Repository
}

func NewService(profileRepo *db.Repository) Service {
	return Service{repository: profileRepo}
}

var nextAvailableID = 1

func promoteNextAvailableID() {
	nextAvailableID++
}

func (s *Service) IsUserInDB(id int) bool {
	if s.repository.IsUserInDB(id) {
		return true
	}
	return false
}

func (s *Service) UpdateUserProfile(userID int, newProfile db.UserProfile) {
	s.repository.UpdateProfile(userID, newProfile)
}

func (s *Service) CreateNewProfile(newProfile db.UserProfile) {
	// Add the new profile to the slice.
	s.repository.NewProfile(nextAvailableID, newProfile)
	promoteNextAvailableID()
}

func (s *Service) GetProfileByID(id int) db.UserProfile {
	return s.repository.GetProfileByID(id)
}
