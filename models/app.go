package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	User_id    uuid.UUID      `gorm:"column:user_id; primaryKey;type:uuid;default:uuid_generate_v4()" json:"user_id" validate:"omitempty,uuid4"`
	Username   string         `gorm:"column:username; unique" json:"username" validate:"required" `
	Email      string         `gorm:"column:email; unique" json:"email" validate:"required" `
	Password   string         `gorm:"column:password" json:"password" validate:"required"`
	Is_Admin   bool           `gorm:"column:is_admin; default:false" json:"is_admin"`
	Created_at time.Time      `gorm:"column:created_at" json:"-"`
	Updated_at time.Time      `gorm:"column:updated_at" json:"-"`
	Deleted_at gorm.DeletedAt `gorm:"column:deleted_at; index" json:"-"`
}

type Post struct {
	Post_id    uuid.UUID      `gorm:"column:post_id;primaryKey;type:uuid;default:uuid_generate_v4()" json:"post_id" `
	Title      string         `gorm:"column:title" json:"title"  validate:"required"`
	Body       string         `gorm:"column:body" json:"body"  validate:"required"`
	User_id    uuid.UUID      `gorm:"column:user_id;foreignKey" json:"user_id" validate:"omitempty,uuid4"`
	User       User           `gorm:"references:User_id;" json:"-"  validate:"omitempty,uuid4"`
	Commands   []Command      `gorm:"-" json:"commands"`
	Category   []Category     `gorm:"-" json:"category"`
	Categories pq.StringArray `gorm:"column:categories; type:text[]" json:"-"`
	Created_at time.Time      `gorm:"column:created_at" json:"-"`
	Updated_at time.Time      `gorm:"column:updated_at" json:"-"`
	Deleted_at gorm.DeletedAt `gorm:"column:deleted_at; index" json:"-"`
}

type Command struct {
	Command_id uuid.UUID      `gorm:"column:command_id;primaryKey;type:uuid;default:uuid_generate_v4()" json:"command_id"`
	Content    string         `gorm:"column:content" json:"content" validate:"required"`
	User_id    uuid.UUID      `gorm:"column:user_id;foreignKey" json:"user_id"`
	User       User           `gorm:"references:User_id" json:"-" validate:"omitempty"`
	Post_id    uuid.UUID      `gorm:"column:post_id;foreignKey" json:"post_id"`
	Post       Post           `gorm:"references:Post_id" json:"-" validate:"omitempty"`
	UserName   string         `gorm:"column:username" json:"username"`
	Created_at time.Time      `gorm:"column:created_at" json:"-"`
	Updated_at time.Time      `gorm:"column:updated_at" json:"-"`
	Deleted_at gorm.DeletedAt `gorm:"column:deleted_at; index" json:"-"`
}

type Category struct {
	Category_id   uuid.UUID      `gorm:"column:category_id;primaryKey;type:uuid;default:uuid_generate_v4()" json:"category_id"`
	Category_Name string         `gorm:"category_name" json:"category_name" validate:"required"`
	Description   string         `gorm:"description" json:"description"  validate:"required"`
	Created_at    time.Time      `gorm:"column:created_at" json:"-"`
	Updated_at    time.Time      `gorm:"column:updated_at" json:"-"`
	Deleted_at    gorm.DeletedAt `gorm:"column:deleted_at; index" json:"-"`
}

type Login struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
