package controller

import (
	swgModels "github.com/cfif1982/cards/internal/infrastructure/swagger/models"
	"github.com/google/uuid"
)

func (c *controller) GetUserByID(userID string) (*swgModels.User, error) {
	// конвертируем строку в uuid
	userUUID, _ := uuid.Parse(userID)

	user, err := c.userUseCases.GetByID(userUUID)

	if err != nil {
		return nil, err
	}

	// вернуть нужно структуру банка для ответа
	// формируем структуру для ответа. С ней же будем рабоать для вставки в бд
	result := swgModels.User{
		ID:        user.ID.String(),
		Name:      user.Name,
		LastName:  user.LastName,
		Email:     user.Email,
		Telephone: user.Telephone,
		Login:     user.Login,
		Password:  user.Password,
	}

	return &result, nil
}
