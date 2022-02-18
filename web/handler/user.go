package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/masibw/go-docker-template/domain/entity"
	"github.com/masibw/go-docker-template/infrastructure/log"
	"github.com/masibw/go-docker-template/usecase"
	"net/http"
)

type UserHandler struct {
	userUC *usecase.UserUseCase
}

func NewUserHandler(userUC *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUC: userUC,
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	log.Logger.Debugf(id)
	user, err := h.userUC.GetUser(id)
	if err != nil {
		log.Logger.Errorf("%v", errors.Is(err, entity.ErrUserNotFound))
		if errors.Is(err, entity.ErrUserNotFound) {
			c.JSON(http.StatusBadRequest, entity.ErrUserNotFound.Error())
			return
		}
		log.Logger.Errorf(err.Error())
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
	c.JSON(http.StatusOK, user)
}
