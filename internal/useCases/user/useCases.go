package user

import (
	"log/slog"

	userRepo "github.com/cfif1982/cards/internal/repositories/user"
	"github.com/cfif1982/cards/internal/useCases/models"
)

type UseCases interface {
	Add(name, lastName, email, telephone, login, password string) (*models.User, error)
	GetByID(userID int64) (*models.User, error)
}

type userUseCases struct {
	log      *slog.Logger
	userRepo userRepo.UserRepo
}

func NewUseCases(log *slog.Logger, userRepo userRepo.UserRepo) UseCases {
	return &userUseCases{
		log:      log,
		userRepo: userRepo,
	}
}
