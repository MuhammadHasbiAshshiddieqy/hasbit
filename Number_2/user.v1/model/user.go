package model

import (
	"errors"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

// User - struct for user
type User struct{
	ID			uuid.UUID	`json:"id"`
	Name		string		`json:"name"`
	Email		string		`json:"email"`
	Password	string		`json:"password"`
}

// BeforeCreate - exec before create new user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
  
	if !u.isValid() {
	  err = errors.New("can't save invalid data")
	}
	return
}

func (u User) isValid() bool {
	if u.ID == uuid.Nil {
		return false
	}
	if u.Name == "" {
		return false
	}
	if u.Email == "" {
		return false
	}
	if u.Password == "" {
		return false
	}
	return true
}