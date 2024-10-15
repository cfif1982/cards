package controller

import (
	swgModels "github.com/cfif1982/cards/internal/infrastructure/swagger/models"
	"github.com/google/uuid"
)

func (c *controller) GetBankByID(bankID string) (*swgModels.Bank, error) {
	// конвертируем строку в uuid
	bankUUID, _ := uuid.Parse(bankID)

	bank, err := c.bankUseCases.GetByID(bankUUID)

	if err != nil {
		return nil, err
	}

	// вернуть нужно структуру банка для ответа
	// формируем структуру для ответа. С ней же будем рабоать для вставки в бд
	result := swgModels.Bank{
		ID:        bank.ID.String(),
		Address:   bank.Address,
		Bik:       bank.BIK,
		Name:      bank.Name,
		Telephone: bank.Telephone,
	}

	return &result, nil
}
