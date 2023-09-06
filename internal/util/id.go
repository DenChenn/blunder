package util

import (
	"github.com/google/uuid"
)

// GenerateRandomId returns a random uuid string
func GenerateRandomId() string {
	id, _ := uuid.NewUUID()
	return id.String()
}
