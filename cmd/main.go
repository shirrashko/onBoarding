package main

import (
	"github.com/shirrashko/BuildingAServer-step2/pkg/api"
)

func main() {
	server := api.NewServer(api.Router)
	server.ListenAndServe()
}

//todo: check if I need to include in this step the health check, and if so, where to include it in the code
