package db

import (
	"database/sql"
	"fmt"
	"github.com/shirrashko/BuildingAServer-step2/pkg/api/model"
)

// ProfileRepository This ProfileRepository struct will encapsulate the operations related to the user profiles using the PostgreSQL database connection.
type ProfileRepository struct {
	client *sql.DB
}

func NewProfileRepository(client *sql.DB) ProfileRepository {
	return ProfileRepository{client: client}
}

func NewDbClient() (*sql.DB, error) {
	// Create a new *sql.DB instance
	connStr := "user=srashkovits dbname=postgres host=localhost port=6432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Error opening database connection: %v\n", err)
		return nil, err
	}
	defer func(db *sql.DB) {
		db.Close()
	}(db) // Close the database connection when the program exits
	return db, nil
}

// implementation of the methods of the ProfileRepository object, which regard to the db contains users profile info

func (repo *ProfileRepository) IsUserInDB(id int) bool {
	query := "SELECT id FROM userProfiles WHERE id = $1" //todo: understand how to put id instead of 1
	var userID int
	err := repo.client.QueryRow(query, id).Scan(&userID)
	if err == nil {
		return true // User with the given id was found
	} else if err != sql.ErrNoRows {
		fmt.Printf("Error checking user existence: %v\n", err)
	}
	return false // User with the given id was not found
}

func (repo *ProfileRepository) UpdateProfile(userID int, newProfile model.UserProfile) error {
	query := "UPDATE userProfiles SET username = $1, full_name = $2, bio = $3, profile_pic_url = $4 WHERE id = $5"
	_, err := repo.client.Exec(query, newProfile.Username, newProfile.FullName, newProfile.Bio, newProfile.ProfilePicURL, userID)
	if err != nil {
		fmt.Printf("Error updating profile: %v\n", err)
		return err
	}
	return nil
}

func (repo *ProfileRepository) NewProfile(userID int, newProfile model.UserProfile) {
	repo.client[userID] = newProfile
}

func (repo *ProfileRepository) GetProfileByID(id int) model.UserProfile {
	return repo.client[id]
}
