package main

import (
	"github.com/masibw/go-docker-template/infrastructure/log"
	"github.com/masibw/go-docker-template/infrastructure/repository"
	"github.com/masibw/go-docker-template/usecase"
	"github.com/masibw/go-docker-template/web"
	"github.com/masibw/go-docker-template/web/handler"
)

func main() {

	userRepo := repository.NewUserRepository()
	userUC := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUC)
	e := web.NewServer(userHandler)

	if err := e.Run(":8080"); err != nil {
		if err != nil {
			log.Logger.Fatalf(err.Error())
		}
	}
}
