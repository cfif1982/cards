package controller

import (
	"cards/internal/infrastructure/repositories/bank"
	respBank "cards/internal/infrastructure/swagger/models"
	"log/slog"
)

type controller struct {
	log      *slog.Logger
	bankRepo bank.BankRepo
}

type Controller interface {
	AddBank(
		name, address, telephone string,
		bik uint32,
	) (*respBank.Bank, error)
}

func NewController(log *slog.Logger, bankRepo bank.BankRepo) Controller {
	return &controller{
		log:      log,
		bankRepo: bankRepo,
	}
}
