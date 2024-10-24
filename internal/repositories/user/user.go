package user

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"

	"github.com/cfif1982/cards/internal/config"
	"github.com/cfif1982/cards/internal/models"
)

const dbConnFormat = "host=%s user=%s password=%s dbname=%s sslmode=disable"

type userRepo struct {
	log *slog.Logger
	db  *sql.DB
}

type UserRepo interface {
	AddUser(user *models.User) error
	GetUserByID(userID uuid.UUID) (*models.User, error)
}

func NewUserRepo(log *slog.Logger, cfg *config.Config) (UserRepo, error) {
	// DSN для СУБД
	// Не стал делать общую БД для всех репозиториев, т.к. я могу захотеть для разных репозиториев использолвать разные БД
	// Поэтому саму БД храню в репозитории
	databaseDSN := fmt.Sprintf(dbConnFormat, cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)

	db, err := sql.Open("pgx", databaseDSN)
	if err != nil {
		return nil, err
	}

	// начинаю миграцию
	// Т.к. делаю миграцию, то не нужно пинговать базу
	log.Info("Start migrating user database")

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

	log.Info("migrating user database finished")

	return &userRepo{
		log: log,
		db:  db,
	}, nil
}
