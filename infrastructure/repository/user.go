package repository

import "github.com/masibw/go-docker-template/domain/entity"

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindByID(id string) (*entity.User, error) {
	user := &entity.User{
		ID:   "example-id",
		Name: "example-name",
	}
	if id == user.ID {
		return user, nil
	}
	return nil, entity.ErrUserNotFound
}
