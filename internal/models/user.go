package models

import (
	swgModels "github.com/cfif1982/cards/internal/infrastructure/swagger/models"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	LastName  string
	Email     string
	Login     string
	Password  string
	Telephone string
}

// DTO для передачи юзера между слоями
//************************************

// DTO из handler в domain
type NewUser struct {
	Name      string
	LastName  string
	Email     string
	Login     string
	Password  string
	Telephone string
}

// конвертируем модель swagger.NewUser в доменную модель NewUser
func ConvertSwaggerToDomainNewUser(swgNewUser *swgModels.NewUser) *NewUser {
	newUser := NewUser{
		Name:      swgNewUser.FirstName,
		LastName:  swgNewUser.LastName,
		Email:     swgNewUser.Email,
		Login:     swgNewUser.Login,
		Password:  swgNewUser.Password,
		Telephone: swgNewUser.Telephone,
	}

	return &newUser
}
