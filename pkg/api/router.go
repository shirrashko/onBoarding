package api

import (
	"github.com/shirrashko/BuildingAServer-step2/cmd/config"
	healthAPI "github.com/shirrashko/BuildingAServer-step2/pkg/api/health"
	profileAPI "github.com/shirrashko/BuildingAServer-step2/pkg/api/profile"
	healthBL "github.com/shirrashko/BuildingAServer-step2/pkg/bl/health"
	profileBL "github.com/shirrashko/BuildingAServer-step2/pkg/bl/profile"
	"github.com/shirrashko/BuildingAServer-step2/pkg/db"
	healthDB "github.com/shirrashko/BuildingAServer-step2/pkg/repository/health"
	profileDB "github.com/shirrashko/BuildingAServer-step2/pkg/repository/profile"
)

type Handlers struct {
	handlers []handler
}

func Router(conf config.Config) (Handlers, error) {
	// Hierarchy chain: handler-> service -> repo -> clientDB

	dbClient, err := db.NewDBClient(conf.DBConfig)
	if err != nil {
		return Handlers{}, err
	}

	// profile
	profileRepo := profileDB.NewRepository(dbClient)
	profileService := profileBL.NewService(&profileRepo)
	profileHandler := profileAPI.NewHandler(&profileService)

	// health
	healthRepo := healthDB.NewRepository(dbClient)
	healthService := healthBL.NewService(&healthRepo)
	healthHandler := healthAPI.NewHandler(&healthService)

	return Handlers{handlers: []handler{profileHandler, healthHandler}}, nil
}
