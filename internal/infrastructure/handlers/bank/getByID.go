package bank

import (
	"github.com/go-openapi/runtime/middleware" // для хэндлера нужна

	"github.com/cfif1982/cards/internal/infrastructure/swagger/restapi/operations/bank"
)

// Хэндлер должен иметь вид: func(params bank.AddBankParams, principal interface{}) middleware.Responder
// здесь params bank.AddBankParams - параметры запроса.
// В Swagger-сгенерированных хэндлерах, параметр principal типа interface{} часто используется для передачи информации о текущем пользователе или клиенте,
// который выполняет запрос
func (h *bankHandlers) GetById(params bank.GetBankByUUIDParams, principal interface{}) middleware.Responder {
	// в yaml описано, что данному запросу поступает в теле данные вида NewBank:
	// - description: Create a new bank in the base
	//         in: body
	//         name: body - вот тут можно изменить имя переменной
	// поэтому в теле запроса передается структура models.NewBank.
	// К полям этой модели можно обращаться через params.Body

	// получаем данные из запроса
	bankID := params.BankID

	// Если нет имени банка, то возвращаем код 400: Invalid input
	if bankID == 0 {
		return bank.NewGetBankByUUIDBadRequest()
	}

	// вызываем метод из контроллера, т.к. мы не можем напрямую работать с репозиторием
	//*********************************************************************************
	result, err := h.controller.GetBankByID(bankID)

	// TODO: пока не сделал нормальный код на этот случай
	if err != nil {
		return bank.NewAddBankUnprocessableEntity()
	}

	// структура ответа:
	// возвращаем код ответа и через .WithPayload() результат. Тип результата смотрим во всплывающей подсказке
	// клиент получит в ответе  JSON-объект с информацией о созданном банке, которая описана в models.Bank
	return bank.NewAddBankOK().WithPayload(result)

}
