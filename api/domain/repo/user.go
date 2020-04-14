package repo

import "github.com/skanehira/sample-web-app/api/domain/model"

type User interface {
	Users() ([]model.User, error)
	User(id int) (*model.User, error)
}
