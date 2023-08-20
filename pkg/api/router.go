package api

import (
	"github.com/shirrashko/BuildingAServer-step2/pkg/api/profile"
	"github.com/shirrashko/BuildingAServer-step2/pkg/bl"
	"github.com/shirrashko/BuildingAServer-step2/pkg/db"
)

type Handlers struct {
	handlers []profile.Handler
}

//createTableSQL := `
//		CREATE TABLE IF NOT EXISTS userProfiles (
//			id SERIAL PRIMARY KEY,
//			username VARCHAR(50) NOT NULL UNIQUE,
//			full_name VARCHAR(100),
//			bio TEXT,
//			profile_pic_url VARCHAR(200)
//		);
//	`

func Router() (Handlers, error) {
	// chain: handler->service->repo->clientDB
	dbClient, err := db.NewDbClient() // need to send an object of type *sql.DB.
	if err != nil {
		return Handlers{}, err
	}
	profileRepo := db.NewProfileRepository(dbClient)
	profileService := bl.NewService(&profileRepo)
	profileHandler := profile.NewHandler(&profileService)
	return Handlers{handlers: []profile.Handler{profileHandler}}, nil
}
