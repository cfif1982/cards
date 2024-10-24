package bank

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/cfif1982/cards/internal/models"
)

// добавить банк
func (b *bankRepo) GetBankByID(bankID uuid.UUID) (*models.Bank, error) {
	// настраиваем squirrel для работы с postgres
	psq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	// создаем запрос на добавление банка
	query, args, _ := psq.
		Select("id", "address", "bik", "name", "telephone").From("banks").Where(sq.Eq{"id": bankID.String()}).
		ToSql()

	// создаю контекст для запроса
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// выполняю запрос
	row := b.db.QueryRowContext(ctx, query, args)

	// в эту переменную будет сканиться результат запроса
	var id uuid.UUID
	var bik uint32
	var address, name, telephone string

	err := row.Scan(&id, &address, &bik, &name, &telephone)

	if err != nil {
		return nil, err
	}

	bank := models.Bank{
		ID:        id,
		Name:      name,
		Address:   address,
		BIK:       bik,
		Telephone: telephone,
	}

	return &bank, nil
}
