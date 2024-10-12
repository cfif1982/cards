package bootstrap

import (
	"log"
	"log/slog"

	"github.com/go-openapi/loads"

	"github.com/cfif1982/cards/internal/infrastructure/swagger/restapi"
	"github.com/cfif1982/cards/internal/infrastructure/swagger/restapi/operations"
	"github.com/cfif1982/cards/internal/infrastructure/swagger/restapi/operations/bank"

	"github.com/cfif1982/cards/internal/config"
	"github.com/cfif1982/cards/internal/controller"
	"github.com/cfif1982/cards/internal/infrastructure/handlers"
	bankRepo "github.com/cfif1982/cards/internal/repositories/bank"
	userRepo "github.com/cfif1982/cards/internal/repositories/user"
)

type Bootstraper struct {
	cfg        *config.Config
	log        *slog.Logger
	server     *restapi.Server
	controller controller.Controller
	handlers   *handlers.Handlers
}

func NewBootstraper(cfg *config.Config, log *slog.Logger) *Bootstraper {
	return &Bootstraper{
		cfg: cfg,
		log: log,
	}
}

func (b *Bootstraper) Run() {
	// создаем репозиторий для банка
	bankRepo, err := bankRepo.NewBankRepo(b.log, b.cfg)

	// создаем репозиторий для юзера
	userRepo, err := userRepo.NewUserRepo(b.log, b.cfg)

	if err != nil {
		panic("Bank Repo error: " + err.Error())
	}

	// создаем контроллер
	b.controller = controller.NewController(b.log, bankRepo, userRepo)

	// создаем хэндлер
	b.handlers = handlers.NewHandlers(b.log, b.controller)

	// Создаем swagger сервер
	//*****************************************************
	// Загрузка спецификации Swagger
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	// Создание нового экземпляра API
	api := operations.NewCardsAPI(swaggerSpec)

	// Создание сервера
	b.server = restapi.NewServer(api)

	b.server.Port = b.cfg.Server.Port
	b.server.Host = b.cfg.Server.Host

	// по завершении работы останавливаем сервер
	defer b.server.Shutdown()

	// Привязываем свои хэндлеры к сгенерированным маршрутам
	b.setupHandlers(api)

	// Настройка авторизации через API-ключ
	api.APIKeyAuth = func(token string) (interface{}, error) {
		// тут будет проверка токена
		// if token == "your_secret_api_key" {
		// 	return true, nil
		// }
		// return nil, errors.New(401, "invalid API key")
		return true, nil
	}

	// Настраиваем хэндлеры - тут идут сгенерированные хэндлеры.
	// Если какой-то хэндлер не настроен, то будет действие по-умолчанию - сообщение типа:
	// operation bank.AddBank has not yet been implemented
	b.server.ConfigureAPI()

	// Запуск сервера
	if err = b.server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

// Привязываем свои хэндлеры к сгенерированным маршрутам.
func (b *Bootstraper) setupHandlers(api *operations.CardsAPI) {
	// все хэндлеры в моем случае разделены на группы,
	// т.к. при создании этих обработчиков в yaml я использовал tag
	// поэтому нужно брать конкретный пакет: Bank, Card, User и в нем выбирать нужный хэндлер
	api.BankAddBankHandler = bank.AddBankHandlerFunc(b.handlers.Bank.Add)
}
