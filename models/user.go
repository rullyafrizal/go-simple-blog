package models

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64       `json:"id" gorm:"primaryKey"`
	Name      string       `json:"name" gorm:"size:255"`
	Email     string       `json:"email" gorm:"size:255"`
	Password  string       `json:"-" gorm:"size:255"`
	Posts     []Post       `json:"posts"`
	CreatedAt sql.NullTime `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt sql.NullTime `json:"updated_at" gorm:"autoUpdateTime"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Password != "" {
		u.Password, err = HashPassword(u.Password)
		if err != nil {
			return err
		}
	}
	return nil
}
