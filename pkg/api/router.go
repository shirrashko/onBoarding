package api

import (
	"github.com/shirrashko/BuildingAServer-step2/pkg/api/profile"
	"github.com/shirrashko/BuildingAServer-step2/pkg/bl"
	"github.com/shirrashko/BuildingAServer-step2/pkg/db"
)

type Handlers struct {
	handlers []profile.Handler
}

func Router() Handlers {
	// chain: handler->service->repo->clientDB
	dbClient := db.NewDbClient()
	profileRepo := db.NewProfileRepository(dbClient)
	profileService := bl.NewService(&profileRepo)
	profileHandler := profile.NewHandler(&profileService)
	return Handlers{handlers: []profile.Handler{profileHandler}}
}
