package controller

import (
	"log/slog"

	"github.com/cfif1982/cards/internal/infrastructure/repositories/bank"
	respBank "github.com/cfif1982/cards/internal/infrastructure/swagger/models"
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
