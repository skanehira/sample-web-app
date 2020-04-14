package model

import "time"

type User struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	Password string     `json:"password"`
	Email    string     `json:"email"`
	CreateAt time.Time  `json:"create_at"`
	DeleteAt *time.Time `json:"-"`
}
