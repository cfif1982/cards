package controller

import (
	swgModels "github.com/cfif1982/cards/internal/infrastructure/swagger/models"
)

func (c *controller) GetBankByID(bankID int64) (*swgModels.Bank, error) {
	bank, err := c.bankUseCases.GetByID(bankID)

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
