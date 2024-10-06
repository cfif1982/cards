package bank

import (
	"github.com/google/uuid"

	"github.com/cfif1982/cards/internal/useCases/models"
)

// Создаем новый объект. Эта функция нужна для работы в других пакетах, когда нужно получить данные
// и поместить их в объект (это про DDD архитектуру).
func (b *bankUseCases) Add(name, address, telephone string, bik uint32) (*models.Bank, error) {
	// генерируем uuid - для дальнейшей вставки в БД
	uuid := uuid.New()

	bank := models.Bank{
		UUID:      uuid,
		Name:      name,
		Address:   address,
		BIK:       bik,
		Telephone: telephone,
	}

	// сохраняем его в БД
	err := b.bankRepo.AddBank(&bank)
	if err != nil {
		return nil, err
	}

	return &bank, nil
}
