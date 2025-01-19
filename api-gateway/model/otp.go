package model

import "github.com/google/uuid"

type Otp struct {
	ID  uuid.UUID
	Otp string
}
