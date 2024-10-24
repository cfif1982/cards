package user

import (
	"github.com/cfif1982/cards/internal/models"
	"github.com/google/uuid"
)

// Создаем новый объект. Эта функция нужна для работы в других пакетах, когда нужно получить данные
// и поместить их в объект (это про DDD архитектуру).
func (u *userUseCase) GetByID(userID uuid.UUID) (*models.User, error) {
	// ищем в БД
	user, err := u.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
