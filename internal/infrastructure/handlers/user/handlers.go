package user

import (
	"log/slog"

	"github.com/go-openapi/runtime/middleware"

	"github.com/cfif1982/cards/internal/controller"
	"github.com/cfif1982/cards/internal/infrastructure/swagger/restapi/operations/user"
)

type Handlers interface {
	Add(params user.AddUserParams, principal interface{}) middleware.Responder
	GetById(params user.GetUserByIDParams, principal interface{}) middleware.Responder
}

// закрываем структуру и делаем для взаимодействия с ней интерфейс
// Делаем такой прием для всех объектов из разных слоев
type userHandlers struct {
	log        *slog.Logger
	controller controller.Controller
}

func NewHandlers(log *slog.Logger, controller controller.Controller) Handlers {
	return &userHandlers{
		log:        log,
		controller: controller,
	}
}
