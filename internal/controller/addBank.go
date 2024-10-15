package controller

import (
	swgModels "github.com/cfif1982/cards/internal/infrastructure/swagger/models"
	"github.com/cfif1982/cards/internal/models"
)

func (c *controller) AddBank(data *swgModels.NewBank) (*swgModels.Bank, error) {
	// конвертируем данные из swagger в domain
	newBank := models.ConvertSwaggerToDomainNewBank(data)

	bank, err := c.bankUseCases.Add(newBank)

	if err != nil {
		return nil, err
	}

	// конвертируем банк из структуры domain в структуру swagger для ответа.
	result := swgModels.Bank{
		ID:        bank.ID.String(),
		Address:   bank.Address,
		Bik:       bank.BIK,
		Name:      bank.Name,
		Telephone: bank.Telephone,
	}

	return &result, nil
}
