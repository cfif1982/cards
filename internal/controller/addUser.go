package controller

import (
	swgModels "github.com/cfif1982/cards/internal/infrastructure/swagger/models"
	"github.com/cfif1982/cards/internal/models"
)

func (c *controller) AddUser(data *swgModels.NewUser) (*swgModels.User, error) {
	// конвертируем данные из swagger в domain
	newUser := models.ConvertSwaggerToDomainNewUser(data)

	user, err := c.userUseCases.Add(newUser)

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
