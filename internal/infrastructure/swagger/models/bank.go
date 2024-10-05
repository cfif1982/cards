// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Bank bank
//
// swagger:model Bank
type Bank struct {

	// address
	// Example: Moscow, Kremlin
	Address string `json:"address,omitempty"`

	// bik
	// Example: 44525974
	Bik uint32 `json:"bik,omitempty"`

	// name
	// Example: Main Bank
	Name string `json:"name,omitempty"`

	// telephone
	// Example: 565-56-56
	Telephone string `json:"telephone,omitempty"`

	// uuid
	// Example: 53aa35c8-e659-44b2-882f-f6056e443c99
	UUID int64 `json:"uuid,omitempty"`
}

// Validate validates this bank
func (m *Bank) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this bank based on context it is used
func (m *Bank) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Bank) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Bank) UnmarshalBinary(b []byte) error {
	var res Bank
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}