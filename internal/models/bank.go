package models

import (
	swgModels "github.com/cfif1982/cards/internal/infrastructure/swagger/models"
	"github.com/google/uuid"
)

type Bank struct {
	ID        uuid.UUID
	Name      string
	Address   string
	BIK       uint32
	Telephone string
}

// DTO для передачи банка между слоями
//************************************

// DTO из handler в domain
type NewBank struct {
	Name      string
	Address   string
	BIK       uint32
	Telephone string
}

// конвертируем модель swagger.NewBank в доменную модель NewBank
func ConvertSwaggerToDomainNewBank(swgNewBank *swgModels.NewBank) *NewBank {
	newBank := NewBank{
		Name:      swgNewBank.Name,
		Address:   swgNewBank.Address,
		BIK:       swgNewBank.Bik,
		Telephone: swgNewBank.Telephone,
	}

	return &newBank
}
