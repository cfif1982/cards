package handlers

import (
	"log/slog"

	"github.com/cfif1982/cards/internal/controller"
	"github.com/cfif1982/cards/internal/infrastructure/handlers/bank"
)

type Handlers struct {
	log  *slog.Logger
	Bank bank.Handlers
	// TODO: добавить хэндлеры для User и Card
}

func NewHandlers(log *slog.Logger, controller controller.Controller) *Handlers {
	return &Handlers{
		log:  log,
		Bank: bank.NewHandlers(log, controller),
	}
}