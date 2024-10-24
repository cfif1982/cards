# генерирем go-файлы сервера swagger

# Путь к сгенерированным файлам
SWAGGER_DIR = .\internal\infrastructure\swagger

# Команда для создания каталога
MKDIR_CMD = if not exist $(SWAGGER_DIR) mkdir $(SWAGGER_DIR)

# Очистка сгенерированных файлов swagger
clean_swagger:
	@echo Swagger clean
	if exist $(SWAGGER_DIR) (rd /S /Q $(SWAGGER_DIR)) else (echo Directory does not exist or is empty.)

# Компиляция swagger
swagger:
	@echo swagger compilation
	$(MKDIR_CMD)
	swagger generate server --with-flatten=full --target=$(SWAGGER_DIR)

# Компиляция моков
REPO_DIR = github.com/cfif1982/cards/internal/repositories
MOCK_DIR = ./internal/repositories/mock
MOCK_DIR_WIN = .\internal\repositories\mock # для удаления файлов командой del на windows нужны правильные слеши

# Очистка сгенерированных файлов mock
clean_mock:
	@echo mock clean
	del /Q $(MOCK_DIR_WIN)\mockBank.go
	del /Q $($(MOCK_DIR_WIN))\mockUser.go

# Компиляция mock
mock:
	@echo mock compilation
	mockgen -destination=$(MOCK_DIR)/mockBank.go -package=mocks $(REPO_DIR)/bank BankRepo
	mockgen -destination=$(MOCK_DIR)/mockUser.go -package=mocks $(REPO_DIR)/user UserRepo

# очистка кэша тестов
clean_test_cash:
	go clean -testcache

# запуск тестов
test: clean_test_cash	
	go test ./...

# запуск тестов хэндлеров
test_handlers: clean_test_cash	
	go test "github.com/cfif1982/cards/internal/infrastructure/handlers/..."

# Цель по умолчанию
all: swagger
