package database

import "github.com/gasparvini/internal/entity"

type UserInterface interface {
	Create(user *entity.User) (entity.User, error)
	FindByEmail(emaild string) (*entity.User, error)
}
