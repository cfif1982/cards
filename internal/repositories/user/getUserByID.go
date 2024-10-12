package user

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/cfif1982/cards/internal/useCases/models"
)

// добавить банк
func (u *userRepo) GetUserByID(userID int64) (*models.User, error) {
	// настраиваем squirrel для работы с postgres
	psq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	// создаем запрос на добавление банка
	query, args, _ := psq.
		Select("uuid", "name", "last_name", "email", "telephone", "login", "password").From("users").Where(sq.Eq{"uuid": userID}).
		ToSql()

	// создаю контекст для запроса
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// выполняю запрос
	row := u.db.QueryRowContext(ctx, query, args)

	// в эту переменную будет сканиться результат запроса
	var uuid uuid.UUID
	var name, lastName, email, telephone, login, password string

	err := row.Scan(&uuid, &name, &lastName, &email, &telephone, &login, &password)

	if err != nil {
		return nil, err
	}

	user := models.User{
		UUID:      uuid,
		Name:      name,
		LastName:  lastName,
		Email:     email,
		Telephone: telephone,
		Login:     login,
		Password:  password,
	}

	return &user, nil
}
