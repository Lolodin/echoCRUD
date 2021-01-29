package repository

import "echoServise/model"

type UserStore interface {
	GetUser(id int) *model.User
	GetUsers() []model.User
	RemoveUser(id int) error
	AddUser(name string) error
	UpdateUser(user model.User) error

}