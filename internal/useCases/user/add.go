package user

import (
	"github.com/google/uuid"

	"github.com/cfif1982/cards/internal/useCases/models"
)

// Создаем новый объект. Эта функция нужна для работы в других пакетах, когда нужно получить данные
// и поместить их в объект (это про DDD архитектуру).
func (u *userUseCases) Add(name, lastName, email, telephone, login, password string) (*models.User, error) {
	// генерируем uuid - для дальнейшей вставки в БД
	uuid := uuid.New()

	user := models.User{
		UUID:      uuid,
		Name:      name,
		LastName:  lastName,
		Email:     email,
		Telephone: telephone,
		Login:     login,
		Password:  password,
	}

	// сохраняем его в БД
	err := u.userRepo.AddUser(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
