package controller

import (
	"log/slog"

	swgModels "github.com/cfif1982/cards/internal/infrastructure/swagger/models"
	bankRepo "github.com/cfif1982/cards/internal/repositories/bank"
	"github.com/cfif1982/cards/internal/useCases/bank"
)

type controller struct {
	log          *slog.Logger
	bankUseCases bank.UseCases
}

type Controller interface {
	AddBank(data *swgModels.NewBank) (*swgModels.Bank, error)
}

func NewController(log *slog.Logger, bankRepo bankRepo.BankRepo) Controller {
	return &controller{
		log:          log,
		bankUseCases: bank.NewUseCases(log, bankRepo),
	}
}
