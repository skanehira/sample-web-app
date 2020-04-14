package infra

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/skanehira/sample-web-app/api/domain/model"
	"github.com/skanehira/sample-web-app/api/domain/repo"
)

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) repo.User {
	return &user{db: db}
}

func (u *user) Users() ([]model.User, error) {
	return nil, nil
}

func (u *user) User(id int) (*model.User, error) {
	user := model.User{
		ID: id,
	}

	if err := u.db.Find(&user).Error; err != nil {
		err = errors.Wrap(err, fmt.Sprintf("cannot get user, id=%d", id))
		log.Println(err)
		return nil, err
	}

	return &user, nil
}
