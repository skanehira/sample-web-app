package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/skanehira/sample-web-app/api/domain/service"
)

type UserHandler interface {
	Users(c echo.Context) error
	User(c echo.Context) error
}

type userHandler struct {
	user service.User
}

func NewUserHandler(user service.User) UserHandler {
	return &userHandler{user: user}
}

func (u *userHandler) Users(c echo.Context) error {
	users, err := u.user.Users()
	if err != nil {
		return nil
	}

	return c.JSON(http.StatusOK, users)
}

func (u *userHandler) User(c echo.Context) error {
	paramID := c.QueryParam("id")
	if paramID == "" {
		err := errors.Wrap(ErrMissingUserID, "bad request")
		log.Println(err)
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err))
	}

	id, err := strconv.Atoi(paramID)
	if err != nil {
		err = errors.Wrap(err, "invalid user id")
		log.Println(err)
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err))
	}

	user, err := u.user.User(id)
	if err != nil {
		err = errors.Wrap(err, "cannot get user")
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}
