package repository

import "github.com/masibw/go-docker-template/domain/entity"

type User interface {
	FindByID(id string) (*entity.User, error)
}
