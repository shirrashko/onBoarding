package bl

import (
	"github.com/shirrashko/BuildingAServer-step2/pkg/api/model"
	"github.com/shirrashko/BuildingAServer-step2/pkg/db"
)

// Service The service has a repository with a client field, which is the connection to the database we are working with
type Service struct {
	repository *db.ProfileRepository
}

func NewService(profileRepo *db.ProfileRepository) Service {
	return Service{repository: profileRepo}
}

func (s *Service) IsUserInDB(id int) bool {
	if s.repository.IsUserInDB(id) {
		return true
	}
	return false
}

func (s *Service) UpdateUserProfile(userID int, newProfile model.UserProfile) error {
	return s.repository.UpdateProfile(userID, newProfile)
}

func (s *Service) CreateNewProfile(newProfile model.UserProfile) (int, error) {
	// Add the new profile to the slice.
	newID, err := s.repository.CreateNewProfile(newProfile)
	return newID, err
}

func (s *Service) GetProfileByID(id int) (model.UserProfile, error) {
	return s.repository.GetProfileByID(id)
}

func HealthChecking() bool {
	// In the real world the health-check function will also check connections to other resources that the server
	//depends on.
	return true
}
