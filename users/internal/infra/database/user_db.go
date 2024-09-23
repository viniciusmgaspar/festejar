package database

import (
	"gorm.io/gorm"

	"github.com/gasparvini/internal/entity"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entity.User) (entity.User, error) {
	err := u.DB.Create(user)
	if err != nil {
		return entity.User{}, err.Error
	}

	return entity.User{Name: user.Name, Email: user.Name, Phone: user.Phone}, nil
}

func (u *User) FindByID(id string) (*entity.User, error) {
	var user entity.User
	err := u.DB.First(&user, "id = ?", id).Error

	return &user, err
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.DB.First(&user, "email = ?", email).Error

	return &user, err
}

func (u *User) FindAll() ([]entity.User, error) {
	var users []entity.User

	err := u.DB.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}
