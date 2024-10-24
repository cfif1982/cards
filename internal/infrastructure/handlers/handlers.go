package handlers

import (
	"log/slog"

	"github.com/cfif1982/cards/internal/controller"
	"github.com/cfif1982/cards/internal/infrastructure/handlers/bank"
	"github.com/cfif1982/cards/internal/infrastructure/handlers/user"
)

type Handlers struct {
	log  *slog.Logger
	Bank bank.Handlers
	User user.Handlers
	// TODO: добавить хэндлеры для Card
}

func NewHandlers(log *slog.Logger, controller controller.Controller) *Handlers {
	return &Handlers{
		log:  log,
		Bank: bank.NewHandlers(log, controller),
		User: user.NewHandlers(log, controller),
	}
}
