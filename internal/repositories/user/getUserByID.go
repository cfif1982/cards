package user

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/cfif1982/cards/internal/models"
)

// добавить банк
func (u *userRepo) GetUserByID(userID uuid.UUID) (*models.User, error) {
	// настраиваем squirrel для работы с postgres
	psq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	// создаем запрос на добавление банка
	query, args, _ := psq.
		Select("id", "name", "last_name", "email", "telephone", "login", "password").From("users").Where(sq.Eq{"uuid": userID}).
		ToSql()

	// создаю контекст для запроса
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// выполняю запрос
	row := u.db.QueryRowContext(ctx, query, args)

	// в эту переменную будет сканиться результат запроса
	var id uuid.UUID
	var name, lastName, email, telephone, login, password string

	err := row.Scan(&id, &name, &lastName, &email, &telephone, &login, &password)

	if err != nil {
		return nil, err
	}

	user := models.User{
		ID:        id,
		Name:      name,
		LastName:  lastName,
		Email:     email,
		Telephone: telephone,
		Login:     login,
		Password:  password,
	}

	return &user, nil
}
