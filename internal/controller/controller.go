package controller

import (
	"log/slog"

	swgModels "github.com/cfif1982/cards/internal/infrastructure/swagger/models"
	bankRepo "github.com/cfif1982/cards/internal/repositories/bank"
	userRepo "github.com/cfif1982/cards/internal/repositories/user"
	bankUseCases "github.com/cfif1982/cards/internal/useCases/bank"
	userUseCases "github.com/cfif1982/cards/internal/useCases/user"
)

type controller struct {
	log          *slog.Logger
	bankUseCases bankUseCases.UseCases
	userUseCases userUseCases.UseCases
}

type Controller interface {
	AddBank(data *swgModels.NewBank) (*swgModels.Bank, error)
	GetBankByID(bankID string) (*swgModels.Bank, error)

	AddUser(data *swgModels.NewUser) (*swgModels.User, error)
	GetUserByID(userID string) (*swgModels.User, error)
}

func NewController(log *slog.Logger, bankRepo bankRepo.BankRepo, userRepo userRepo.UserRepo) Controller {
	return &controller{
		log:          log,
		bankUseCases: bankUseCases.NewUseCases(log, bankRepo),
		userUseCases: userUseCases.NewUseCases(log, userRepo),
	}
}
