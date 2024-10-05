package bank

import "github.com/google/uuid"

type Bank struct {
	uuid      uuid.UUID
	name      string
	address   string
	bik       uint32
	telephone string
}

// Создаем новый объект. Эта функция нужна для работы в других пакетах, когда нужно получить данные
// и поместить их в объект (это про DDD архитектуру).
func NewBank(uuid uuid.UUID, name, address, telephone string, bik uint32) *Bank {
	return &Bank{
		uuid:      uuid,
		name:      name,
		address:   address,
		bik:       bik,
		telephone: telephone,
	}
}

// Создаем банк. Проверяем его поля и т.д.
func CreateBank(uuid uuid.UUID, name, address, telephone string, bik uint32) *Bank {
	// проверяем бик, имя, телефон и т.д.

	return NewBank(uuid, name, address, telephone, bik)
}

// возвращщаем поле UUID
func (b *Bank) UUID() uuid.UUID {
	return b.uuid
}

// возвращщаем поле name
func (b *Bank) Name() string {
	return b.name
}

// возвращщаем поле address
func (b *Bank) Address() string {
	return b.address
}

// возвращщаем поле bik
func (b *Bank) BIK() uint32 {
	return b.bik
}

// возвращщаем поле telephone
func (b *Bank) Telephone() string {
	return b.telephone
}
