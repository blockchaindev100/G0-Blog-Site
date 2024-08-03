package models

import "github.com/google/uuid"

type Command struct {
	Command_id uuid.UUID
	Content    string
	User_id    uuid.UUID
	Post_id    uuid.UUID
}
