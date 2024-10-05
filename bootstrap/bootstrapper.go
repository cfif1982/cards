package bootstrap

import (
	"fmt"
	"log"
	"log/slog"

	// эти пакеты нужны для закоменченного когда ниже
	// "os"
	// flags "github.com/jessevdk/go-flags"

	"cards/internal/config"
	"cards/internal/controller"
	bankHandlers "cards/internal/infrastructure/handlers/bank"
	bankRepo "cards/internal/infrastructure/repositories/bank"

	"cards/internal/infrastructure/swagger/restapi"
	"cards/internal/infrastructure/swagger/restapi/operations"
	"cards/internal/infrastructure/swagger/restapi/operations/bank"

	"github.com/go-openapi/loads"
)

const databaseUser = "postgres"
const databasePassword = "123"
const databaseHost = "localhost"
const databaseTable = "cards"

type Bootstraper struct {
	cfg        *config.Config
	log        *slog.Logger
	server     *restapi.Server
	controller controller.Controller
	handlers   *bankHandlers.Handlers
}

func NewBootstraper(cfg *config.Config, log *slog.Logger) *Bootstraper {
	return &Bootstraper{
		cfg: cfg,
		log: log,
	}
}

func (b *Bootstraper) Run() {
	// создаем репозиторий для банка
	// DSN для СУБД
	databaseDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", databaseHost, databaseUser, databasePassword, databaseTable)

	// создаем сам репозиторий для банка
	bankRepo, err := bankRepo.NewBankRepo(b.log, databaseDSN)

	if err != nil {
		panic("Bank Repo error: " + err.Error())
	}

	// создаем контроллер
	b.controller = controller.NewController(b.log, bankRepo)

	// создаем хэндлер
	b.handlers = bankHandlers.NewHandlers(b.log, b.controller)

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

	b.server.Port = b.cfg.Port
	b.server.Host = b.cfg.Host

	// по завершении работы останавливаем сервер
	defer b.server.Shutdown()

	// TODO: не знаю зачем это нужно
	/*
		parser := flags.NewParser(server, flags.Default)
		parser.ShortDescription = "cards"
		parser.LongDescription = "This is a sample User's cards application based on the OpenAPI 2.0 specification."
		server.ConfigureFlags()
		for _, optsGroup := range api.CommandLineOptionsGroups {
			_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
			if err != nil {
				log.Fatalln(err)
			}
		}

		if _, err := parser.Parse(); err != nil {
			code := 1
			if fe, ok := err.(*flags.Error); ok {
				if fe.Type == flags.ErrHelp {
					code = 0
				}
			}
			os.Exit(code)
		}
	*/

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
	api.BankAddBankHandler = bank.AddBankHandlerFunc(b.handlers.AddBankHandler)
}
