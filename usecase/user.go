package usecase

import (
	"fmt"
	"github.com/masibw/go-docker-template/domain/entity"
	"github.com/masibw/go-docker-template/domain/repository"
	"github.com/masibw/go-docker-template/infrastructure/log"
	"github.com/masibw/go-docker-template/usecase/dto"
)

type UserUseCase struct {
	userRepository repository.User
}

func NewUserUseCase(userRepository repository.User) *UserUseCase {
	return &UserUseCase{userRepository: userRepository}
}

func (u *UserUseCase) GetUser(id string) (userDTO *dto.UserDTO, err error) {
	var user *entity.User
	user, err = u.userRepository.FindByID(id)
	if err != nil {
		log.Logger.Debugf("GetUser: %v", err.Error())
		err = fmt.Errorf("get user: %w", err)
		return
	}

	userDTO = dto.NewUserFromEntity(user)
	return
}
