package db

import "github.com/shirrashko/BuildingAServer-step2/pkg/api/model"

type Repository struct {
	client map[int]model.UserProfile // hold all the users profiles, act as the DB. user is access via id
}

var UsersProfiles map[int]model.UserProfile // In this file just for now, later in the exercises will be different

func NewDbClient() map[int]model.UserProfile {
	return UsersProfiles
}

func NewProfileRepository(c map[int]model.UserProfile) Repository {
	return Repository{client: c} //todo: check if map is received to a function as a deep copy or shallow copy
}

// implementation of the methods of the Repository object, which regard to the db contains users profile info

func (repo *Repository) IsUserInDB(id int) bool {
	for key := range repo.client {
		if key == id {
			return true
		}
	}
	return false
}

func (repo *Repository) UpdateProfile(userID int, newProfile model.UserProfile) {
	repo.client[userID] = newProfile
}

func (repo *Repository) NewProfile(userID int, newProfile model.UserProfile) {
	repo.client[userID] = newProfile
}

func (repo *Repository) GetProfileByID(id int) model.UserProfile {
	return repo.client[id]
}
