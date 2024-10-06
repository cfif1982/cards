package controller

import (
	swgModels "github.com/cfif1982/cards/internal/infrastructure/swagger/models"
)

func (c *controller) AddBank(data *swgModels.NewBank) (*swgModels.Bank, error) {
	bank, err := c.bankUseCases.Add(data.Name, data.Address, data.Telephone, data.Bik)

	if err != nil {
		return nil, err
	}

	// вернуть нужно структуру банка для ответа
	// формируем структуру для ответа. С ней же будем рабоать для вставки в бд
	result := swgModels.Bank{
		UUID:      int64(bank.UUID.ID()),
		Address:   bank.Address,
		Bik:       bank.BIK,
		Name:      bank.Name,
		Telephone: bank.Telephone,
	}

	return &result, nil
}
