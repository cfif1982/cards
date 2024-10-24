package bank

import (
	"github.com/google/uuid"

	"github.com/cfif1982/cards/internal/models"
)

// Создаем новый объект. Эта функция нужна для работы в других пакетах, когда нужно получить данные
// и поместить их в объект (это про DDD архитектуру).
func (b *bankUseCase) Add(newBank *models.NewBank) (*models.Bank, error) {
	// генерируем uuid - для дальнейшей вставки в БД
	uuid := uuid.New()

	bank := models.Bank{
		ID:        uuid,
		Name:      newBank.Name,
		Address:   newBank.Address,
		BIK:       newBank.BIK,
		Telephone: newBank.Telephone,
	}

	// сохраняем его в БД
	err := b.bankRepo.AddBank(&bank)
	if err != nil {
		return nil, err
	}

	return &bank, nil
}
