package bank

import (
	"log/slog"

	"github.com/cfif1982/cards/internal/models"
	bankRepo "github.com/cfif1982/cards/internal/repositories/bank"
	"github.com/google/uuid"
)

type UseCases interface {
	Add(newBank *models.NewBank) (*models.Bank, error)
	GetByID(bankID uuid.UUID) (*models.Bank, error)
}

type bankUseCase struct {
	log      *slog.Logger
	bankRepo bankRepo.BankRepo
}

func NewUseCases(log *slog.Logger, bankRepo bankRepo.BankRepo) UseCases {
	return &bankUseCase{
		log:      log,
		bankRepo: bankRepo,
	}
}
