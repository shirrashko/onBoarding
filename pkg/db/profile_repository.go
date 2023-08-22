package db

import (
	"database/sql"
	"fmt"
	"github.com/shirrashko/BuildingAServer-step2/cmd/config"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/shirrashko/BuildingAServer-step2/pkg/api/model"
)

// ProfileRepository This ProfileRepository struct will encapsulate the operations related to the user profiles using the PostgreSQL database connection.
type ProfileRepository struct {
	client *sql.DB
}

func NewProfileRepository(client *sql.DB) ProfileRepository {
	return ProfileRepository{client: client}
}

func NewDBClient(connectionInfo config.DBConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", connectionInfo.Host,
		connectionInfo.Port, connectionInfo.User, connectionInfo.Password, connectionInfo.DBName)
	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		fmt.Printf("Error opening database connection: %v\n", err)
		return nil, err
	}

	// Execute the table creation script
	script, err := os.ReadFile("/Users/srashkovits/repos/onboarding/scheme/create_table.sql")
	if err != nil {
		fmt.Printf("Error reading create_tables.sql: %v\n", err)
		return nil, err
	}

	_, err = db.Exec(string(script)) // execute the script create a table named userProfiles with columns names
	// according to the api.model
	if err != nil {
		fmt.Printf("Error executing create_tables.sql: %v\n", err)
		return nil, err
	}

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
	_, err := repo.client.Exec(query, newProfile.Username, newProfile.FullName, newProfile.Bio, newProfile.ProfilePicURL, userID) // execute the prepared SQL query.
	if err != nil {
		fmt.Printf("Error updating profile: %v\n", err)
		return err
	}
	return nil
}

func (repo *ProfileRepository) CreateNewProfile(newProfile model.UserProfile) (int, error) {
	query := "INSERT INTO userProfiles (username, full_name, bio, profile_pic_url) VALUES ($1, $2, $3, $4) RETURNING id"
	var id int
	err := repo.client.QueryRow(query, newProfile.Username, newProfile.FullName, newProfile.Bio, newProfile.ProfilePicURL).Scan(&id)
	if err != nil {
		fmt.Printf("Error creating profile: %v\n", err)
		return 0, err
	}
	return id, nil
}

func (repo *ProfileRepository) GetProfileByID(id int) (model.UserProfile, error) {
	query := "SELECT id, username, full_name, bio, profile_pic_url FROM userProfiles WHERE id = $1"
	var userProfile model.UserProfile
	err := repo.client.QueryRow(query, id).Scan(&userProfile.ID, &userProfile.Username, &userProfile.FullName, &userProfile.Bio, &userProfile.ProfilePicURL)

	if err != nil {
		if err == sql.ErrNoRows { // Handle the case where no rows were found.
			return userProfile, fmt.Errorf("no user found with ID %d", id)
		}
		fmt.Printf("Error querying user profile: %v\n", err)
		return userProfile, err
	}
	return userProfile, nil
}
