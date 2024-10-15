package bank

import (
	"log/slog"

	"github.com/go-openapi/runtime/middleware"

	"github.com/cfif1982/cards/internal/controller"
	"github.com/cfif1982/cards/internal/infrastructure/swagger/restapi/operations/bank"
)

type Handlers interface {
	Add(params bank.AddBankParams, principal interface{}) middleware.Responder
	GetById(params bank.GetBankByIDParams, principal interface{}) middleware.Responder
}

// закрываем структуру и делаем для взаимодействия с ней интерфейс
// Делаем такой прием для всех объектов из разных слоев
type bankHandlers struct {
	log        *slog.Logger
	controller controller.Controller
}

func NewHandlers(log *slog.Logger, controller controller.Controller) Handlers {
	return &bankHandlers{
		log:        log,
		controller: controller,
	}
}
