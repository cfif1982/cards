package main

import (
	"log/slog"
	"os"

	"github.com/cfif1982/cards/bootstrap"
	"github.com/cfif1982/cards/internal/config"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	// считываем конфигурацию
	cfg := config.MustLoad()

	// инициализируем логгер
	log := setupLogger(cfg.Env)

	log.Info("starting cards...")

	// создаем bootsraper
	bs := bootstrap.NewBootstraper(cfg, log)

	bs.Run() // запуск bootstraper
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

/*
генерировал командой
swagger generate server --with-flatten=full --target=./internal/infrastructure/swagger

Роман генерировал командой
swagger generate server --spec=openapi2.yaml --with-expand --with-flatten=full --target=./internal/infrastructure/ --server-package=./httpserver --api-package=./swaggerapi --model-package=./httpserver/swaggerapi/models --exclude-main


структура проекта
restapi/configure_cards.go - файл конфигурации, который позволяет настраивать хэндлеры для операций. В нем прописана ConfigureAPI()
*/
