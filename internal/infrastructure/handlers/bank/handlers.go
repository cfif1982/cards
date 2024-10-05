package handlers

import (
	"cards/internal/controller"
	"log/slog"
)

// TODO: могу ли я перенести этот файл в папку handlers и сделать его общим для всех хэндлеров в подпаках?

type Handlers struct {
	log        *slog.Logger
	controller controller.Controller
}

func NewHandlers(log *slog.Logger, controller controller.Controller) *Handlers {
	return &Handlers{
		log:        log,
		controller: controller,
	}
}
