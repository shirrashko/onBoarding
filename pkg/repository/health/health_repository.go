package health

import "database/sql"

type HealthRepository struct {
	client *sql.DB
}

func NewHealthRepository(client *sql.DB) HealthRepository {
	return HealthRepository{client: client}
}

func (h HealthRepository) HealthCheck() bool {
	return true
}
