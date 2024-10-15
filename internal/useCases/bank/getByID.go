package bank

import (
	"github.com/cfif1982/cards/internal/models"
	"github.com/google/uuid"
)

// Создаем новый объект. Эта функция нужна для работы в других пакетах, когда нужно получить данные
// и поместить их в объект (это про DDD архитектуру).
func (b *bankUseCase) GetByID(bankID uuid.UUID) (*models.Bank, error) {
	// ищем в БД
	bank, err := b.bankRepo.GetBankByID(bankID)
	if err != nil {
		return nil, err
	}

	return bank, nil
}
