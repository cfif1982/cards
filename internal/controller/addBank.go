package controller

import (
	"cards/internal/domain/bank"
	respBank "cards/internal/infrastructure/swagger/models"

	"github.com/google/uuid"
)

func (c *controller) AddBank(
	name, address, telephone string,
	bik uint32,
) (*respBank.Bank, error) {

	// генерируем uuid - для дальнейшей вставки в БД
	uuid := uuid.New()

	// создаем объект банк
	bank := bank.CreateBank(
		uuid,
		name,
		address,
		telephone,
		bik,
	)

	// сохраняем его в БД
	err := c.bankRepo.AddBank(bank)
	if err != nil {
		return nil, err
	}

	// вернуть нужно структуру банка для ответа
	// формируем структуру для ответа. С ней же будем рабоать для вставки в бд
	result := respBank.Bank{
		UUID:      int64(bank.UUID().ID()),
		Address:   bank.Address(),
		Bik:       bank.BIK(),
		Name:      bank.Name(),
		Telephone: bank.Telephone(),
	}

	return &result, nil
}
