package service

import "github.com/google/uuid"

func GenerateUUIDString() string {
	return uuid.New().String()
}
