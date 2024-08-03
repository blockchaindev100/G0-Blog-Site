package models

import "github.com/google/uuid"

type Post struct {
	Post_id  uuid.UUID
	Title    string
	Body     string
	Category string
	User_id  uuid.UUID
}
