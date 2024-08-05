package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	User_id    uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"user_id"`
	Username   string    `gorm:"unique" json:"username" validate:"required" `
	Email      string    `gorm:"unique" json:"email" validate:"required" `
	Password   string    `json:"password" validate:"required"`
	Is_Admin   bool      `gorm:"default:false" json:"is_admin" validate:"required"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at gorm.DeletedAt `gorm:"index"`
}

type Post struct {
	Post_id    uuid.UUID `json:"post_id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Title      string    `json:"title"  validate:"required"`
	Body       string    `json:"body"  validate:"required"`
	Category   string    `json:"category"  validate:"required"`
	User_id    uuid.UUID `gorm:"foreignKey" json:"user_id"  validate:"required"`
	User       User      `gorm:"references:User_id;"  validate:"required"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at gorm.DeletedAt `gorm:"index"`
}

type Command struct {
	Command_id uuid.UUID `json:"command_id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Content    string    `json:"content" validate:"required"`
	User_id    uuid.UUID `gorm:"foreignKey" json:"user_id" validate:"required"`
	User       User      `gorm:"references:User_id"`
	Post_id    uuid.UUID `gorm:"foreignKey" json:"post_id" validate:"required"`
	Post       Post      `gorm:"references:Post_id"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at gorm.DeletedAt `gorm:"index"`
}

type Category struct {
	Category_id   uuid.UUID `json:"category_id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Category_Name string    `json:"category_name" validate:"required"`
	Description   string    `json:"description"  validate:"required"`
	Created_at    time.Time
	Updated_at    time.Time
	Deleted_at    gorm.DeletedAt `gorm:"index"`
}

type Category_Post_Mapping struct {
	Id          uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Post_id     uuid.UUID `json:"post_id" validate:"required"`
	Category_id uuid.UUID `json:"category_id" validate:"required"`
	Created_at  time.Time
	Updated_at  time.Time
	Deleted_at  gorm.DeletedAt `gorm:"index"`
}

type Login struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
