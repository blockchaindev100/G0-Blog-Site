package models

import (
	"github.com/google/uuid"
)

type User struct {
	User_id  uuid.UUID `json:"user_id" gorm:"primaryKey"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Is_Admin bool      `json:"is_admin"`
}
