package models

import "github.com/google/uuid"

type Bank struct {
	UUID      uuid.UUID
	Name      string
	Address   string
	BIK       uint32
	Telephone string
}
