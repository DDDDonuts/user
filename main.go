package main

import (
	"backend/account/api"
	"backend/account/inmem"
	"backend/account/service"
)

func main() {

	// user repository
	inMemRepo := inmem.NewInMemUserRepository()

	// user service
	svc := service.NewUserService(inMemRepo)

	// user transport (http api)
	userApi := api.NewHTTPTransport(svc)

	userApi.Run()
}
