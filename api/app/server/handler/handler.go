package handler

import "github.com/skanehira/sample-web-app/api/domain/service"

type Handler interface {
	UserHandler
}

type handler struct {
	UserHandler
}

func NewHandler(userService service.User) Handler {
	return &handler{
		UserHandler: NewUserHandler(userService),
	}
}
