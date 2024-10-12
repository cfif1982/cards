package controller

import (
	swgModels "github.com/cfif1982/cards/internal/infrastructure/swagger/models"
)

func (c *controller) AddUser(data *swgModels.NewUser) (*swgModels.User, error) {
	user, err := c.userUseCases.Add(data.FirstName, data.LastName, data.Email, data.Telephone, data.Login, data.Password)

	if err != nil {
		return nil, err
	}

	// вернуть нужно структуру банка для ответа
	// формируем структуру для ответа. С ней же будем рабоать для вставки в бд
	result := swgModels.User{
		UUID:      int64(user.UUID.ID()),
		Name:      user.Name,
		LastName:  user.LastName,
		Email:     user.Email,
		Telephone: user.Telephone,
		Login:     user.Login,
		Password:  user.Password,
	}

	return &result, nil
}
