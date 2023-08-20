package api

import (
	"github.com/shirrashko/BuildingAServer-step2/pkg/api/profile"
	"github.com/shirrashko/BuildingAServer-step2/pkg/bl"
	"github.com/shirrashko/BuildingAServer-step2/pkg/db"
)

type Handlers struct {
	handlers []profile.Handler
}

func Router() (Handlers, error) {
	// chain: handler->service->repo->clientDB
	dbClient, err := db.NewDbClient() // todo: need to send an object of type *sql.DB ?
	if err != nil {
		return Handlers{}, err
	}
	profileRepo := db.NewProfileRepository(dbClient)
	profileService := bl.NewService(profileRepo)
	profileHandler := profile.NewHandler(&profileService)
	return Handlers{handlers: []profile.Handler{profileHandler}}, nil
}
