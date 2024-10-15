package bank

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/cfif1982/cards/internal/controller"
	"github.com/cfif1982/cards/internal/infrastructure/swagger/models"
	"github.com/cfif1982/cards/internal/infrastructure/swagger/restapi/operations/bank"
	mocks "github.com/cfif1982/cards/internal/repositories/mock"
)

type test struct {
	testName        string             // название теста
	requestBody     string             // тело запроса
	handlerParams   bank.AddBankParams // объект для передачи в параметры хэндлера
	mockReturnError error              // возвращаемая моком ошибка
	expectedCode    int                // ожидаемый тестом код ответа хэндлера
}

func TestAdd(t *testing.T) {
	// создаем хэндлер
	handler, mockBankRepo, _ := setupHandler(t)

	// инициализируем данные для тестов
	tests := initTests()

	// запускаем тесты
	for _, tt := range *tests {
		fmt.Println(tt.testName)

		// Настраиваем mock для текущего теста: принимает любые параметры и возвращает нужную ошибку
		mockBankRepo.EXPECT().AddBank(gomock.Any()).Return(tt.mockReturnError)

		// Вызываем хендлер напрямую с параметрами
		resp := handler.Add(tt.handlerParams, nil)

		// Получаем ResponseRecorder
		recorder := getRecorder(resp)

		// Проверяем код ответа
		assert.Equal(t, tt.expectedCode, recorder.Code, "Ожидался код ответа %d", tt.expectedCode)

		// Проверяем тело ответа
		// assert.Equal(t, tt.expectedBody, recorder.Code.String(), "Ответ тела не соответствует ожидаемому для userID %s", tt.userID)
	}
}

// Получаем ResponseRecorder
func getRecorder(resp middleware.Responder) *httptest.ResponseRecorder {
	// Создаем ResponseRecorder для захвата ответа
	recorder := httptest.NewRecorder()

	// Создаем JSON-продюсер
	jsonProducer := runtime.JSONProducer()

	// Вызываем resp с recorder и jsonProducer
	resp.WriteResponse(recorder, jsonProducer)

	return recorder
}

// создаем и настраиваем хэндлер
func setupHandler(t *testing.T) (Handlers, *mocks.MockBankRepo, *mocks.MockUserRepo) {
	// создаём контроллер
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// создаём объект-заглушку
	mockBankRepo := mocks.NewMockBankRepo(ctrl)
	mockUserRepo := mocks.NewMockUserRepo(ctrl)

	// инициализируем логгер
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	// создаем контроллер
	controller := controller.NewController(log, mockBankRepo, mockUserRepo)

	// создаем хэндлер
	handler := NewHandlers(log, controller)

	return handler, mockBankRepo, mockUserRepo
}

// инициализируем данные для тестов
func initTests() *[]test {
	t := []test{
		{
			testName: "add bank test #1: 200 OK",
			requestBody: `{
				"address": "Moscow, Kremlin",
				"bik": 44525974,
				"name": "Main Bank",
				"telephone": "565-56-56"
			}`,
			handlerParams: bank.AddBankParams{
				Body: &models.NewBank{
					Address:   "Moscow, Kremlin",
					Bik:       44525974,
					Name:      "Main Bank",
					Telephone: "565-56-56",
				},
			},
			mockReturnError: nil,
			expectedCode:    http.StatusOK,
		},
		{
			testName: "add bank test #2: 400 BadRequest",
			requestBody: `{
				"address": "Moscow, Kremlin",
				"bik": 44525974,
				"telephone": "565-56-56"
			}`,
			handlerParams: bank.AddBankParams{
				Body: &models.NewBank{
					Address:   "Moscow, Kremlin",
					Bik:       44525974,
					Telephone: "565-56-56",
				},
			},
			mockReturnError: nil,
			expectedCode:    http.StatusBadRequest,
		},
		{
			testName: "add bank test #3: 422 UnprocessableEntity",
			requestBody: `{
				"address": "Moscow, Kremlin",
				"bik": 44525,
				"name": "Main Bank",
				"telephone": "565-56-56"
			}`,
			handlerParams: bank.AddBankParams{
				Body: &models.NewBank{
					Address:   "Moscow, Kremlin",
					Bik:       44525,
					Name:      "Main Bank",
					Telephone: "565-56-56",
				},
			},
			mockReturnError: nil,
			expectedCode:    http.StatusUnprocessableEntity,
		},
	}

	return &t
}
