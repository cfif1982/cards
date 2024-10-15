package user

import (
	"log/slog"

	"github.com/cfif1982/cards/internal/models"
	userRepo "github.com/cfif1982/cards/internal/repositories/user"
	"github.com/google/uuid"
)

type UseCases interface {
	Add(newUser *models.NewUser) (*models.User, error)
	GetByID(userID uuid.UUID) (*models.User, error)
}

type userUseCase struct {
	log      *slog.Logger
	userRepo userRepo.UserRepo
}

func NewUseCases(log *slog.Logger, userRepo userRepo.UserRepo) UseCases {
	return &userUseCase{
		log:      log,
		userRepo: userRepo,
	}
}
