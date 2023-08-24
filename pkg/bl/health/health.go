package health

import (
	"github.com/shirrashko/BuildingAServer-step2/pkg/repository/health"
)

// Service The service has a repository with a client field, which is the connection to the database we are working with
type Service struct {
	repository *health.HealthRepository
}

func NewService(healthRepo *health.HealthRepository) Service {
	return Service{repository: healthRepo}
}

func (s Service) HealthCheck() bool {
	return s.repository.HealthCheck()
}
