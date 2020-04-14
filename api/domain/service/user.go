package service

import (
	"github.com/skanehira/sample-web-app/api/domain/model"
	"github.com/skanehira/sample-web-app/api/domain/repo"
)

type User interface {
	Users() ([]model.User, error)
	User(id int) (*model.User, error)
}

type user struct {
	repo repo.User
}

func NewUser(r repo.User) User {
	return &user{repo: r}
}

func (u *user) Users() ([]model.User, error) {
	return nil, nil
}

func (u *user) User(id int) (*model.User, error) {
	return u.repo.User(id)
}
