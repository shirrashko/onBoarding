package main

import (
	"github.com/pkg/errors"
	"github.com/shirrashko/BuildingAServer-step2/pkg/api"
)

func main() {
	server, err := api.NewServer(api.Router)
	if err != nil {
		panic(errors.Wrap(err, "Unable to create server"))
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(errors.Wrap(err, "Unable to start the server"))
	}
}
