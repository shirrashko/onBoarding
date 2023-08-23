package api

import (
	"github.com/shirrashko/BuildingAServer-step2/cmd/config"
	"github.com/shirrashko/BuildingAServer-step2/pkg/api/health"
	profileAPI "github.com/shirrashko/BuildingAServer-step2/pkg/api/profile"
	health2 "github.com/shirrashko/BuildingAServer-step2/pkg/bl/health"
	profileBL "github.com/shirrashko/BuildingAServer-step2/pkg/bl/profile"
	healthDB "github.com/shirrashko/BuildingAServer-step2/pkg/db/health"
	profileDB "github.com/shirrashko/BuildingAServer-step2/pkg/db/profile"
)

type Handlers struct {
	handlers []handler
}

func Router(conf config.Config) (Handlers, error) {
	// chain: handler-> service -> repo -> clientDB
	dbClient, err := profileDB.NewDBClient(conf.DBConfig) // todo: need to send an object of type *sql.DB ?
	if err != nil {
		return Handlers{}, err
	}
	profileRepo := profileDB.NewProfileRepository(dbClient)
	profileService := profileBL.NewService(&profileRepo)
	profileHandler := profileAPI.NewHandler(&profileService)

	healthRepo := healthDB.NewHealthRepository(dbClient)
	healthService := health2.NewService(&healthRepo)
	healthHandler := health.NewHandler(&healthService)

	// healthcheckHandler := health.NewHandler()
	return Handlers{handlers: []handler{profileHandler, healthHandler}}, nil
}
