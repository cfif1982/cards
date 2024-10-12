package models

import "github.com/google/uuid"

type User struct {
	UUID      uuid.UUID
	Name      string
	LastName  string
	Email     string
	Login     string
	Password  string
	Telephone string
}
