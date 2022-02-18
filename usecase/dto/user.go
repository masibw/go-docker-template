package dto

import "github.com/masibw/go-docker-template/domain/entity"

type UserDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewUserFromEntity(e *entity.User) *UserDTO {
	return &UserDTO{
		ID:   e.ID,
		Name: e.Name,
	}
}

func (d *UserDTO) ToEntity() *entity.User {
	return &entity.User{
		ID:   d.ID,
		Name: d.Name,
	}
}
