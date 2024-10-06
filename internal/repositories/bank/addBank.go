package bank

import (
	"context"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/cfif1982/cards/internal/useCases/models"
)

// добавить банк
func (b *bankRepo) AddBank(bank *models.Bank) error {
	// настраиваем squirrel для работы с postgres
	psq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	// создаем запрос на добавление банка
	query, args, _ := psq.
		Insert("banks").Columns("uuid", "name", "address", "bik", "telephone").
		Values(bank.UUID, bank.Name, bank.Address, bank.BIK, bank.Telephone).
		ToSql()

	// создаю контекст для запроса
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// выполняю запрос
	_, err := b.db.ExecContext(ctx, query, args...)
	if err != nil {
		// проверяем ошибку на предмет вставки маршрута с названием, которое уже есть в БД
		// создаем объект *pgconn.PgError - в нем будет храниться код ошибки из БД
		var pgErr *pgconn.PgError

		// преобразуем ошибку к типу pgconn.PgError
		if errors.As(err, &pgErr) {
			// если ошибка- запись существует, то возвращаем эту ошибку
			if pgErr.Code == pgerrcode.UniqueViolation {
				return pgErr
			} else {
				return pgErr
			}
		} else {
			return err
		}
	}

	return nil
}
