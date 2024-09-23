package entity

import (
	"time"

	"github.com/gasparvini/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        entity.ID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	Status    string    `gorm:"type:varchar(20);default:'A'" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	Phone     string    `gorm:"type:varchar(20);not null" json:"phone"`
}

func NewUser(name, email, password, phone string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        entity.NewID(),
		Name:      name,
		Email:     email,
		Password:  string(hash),
		Status:    "A",
		Phone:     phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
