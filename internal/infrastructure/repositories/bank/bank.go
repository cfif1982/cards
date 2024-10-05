package bank

import (
	"database/sql"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/cfif1982/cards/internal/domain/bank"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
)

type bankRepo struct {
	log *slog.Logger
	db  *sql.DB
}

type BankRepo interface {
	AddBank(bank *bank.Bank) error
}

func NewBankRepo(log *slog.Logger, databaseDSN string) (BankRepo, error) {
	db, err := sql.Open("pgx", databaseDSN)
	if err != nil {
		return nil, err
	}

	// начинаю миграцию
	// Т.к. делаю миграцию, то не нужно пинговать базу
	log.Info("Start migrating bank database")

	if err := goose.SetDialect("postgres"); err != nil {
		log.Info(err.Error())
	}

	// узнаю текущую папку, чтобы передать путь к папке с миграциями
	ex, err := os.Executable()
	if err != nil {
		log.Info(err.Error())
	}
	exPath := filepath.Dir(ex)

	exPath = exPath + "/migrations"

	err = goose.Up(db, exPath)
	if err != nil {
		log.Info(err.Error() + ": " + exPath)
	}

	log.Info("migrating bank database finished")

	return &bankRepo{
		log: log,
		db:  db,
	}, nil
}
