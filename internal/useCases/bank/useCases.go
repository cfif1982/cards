package bank

import (
	"log/slog"

	bankRepo "github.com/cfif1982/cards/internal/repositories/bank"
	"github.com/cfif1982/cards/internal/useCases/models"
)

type UseCases interface {
	Add(name, address, telephone string, bik uint32) (*models.Bank, error)
	GetByID(bankID int64) (*models.Bank, error)
}

type bankUseCases struct {
	log      *slog.Logger
	bankRepo bankRepo.BankRepo
}

func NewUseCases(log *slog.Logger, bankRepo bankRepo.BankRepo) UseCases {
	return &bankUseCases{
		log:      log,
		bankRepo: bankRepo,
	}
}
