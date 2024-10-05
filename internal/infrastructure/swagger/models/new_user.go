// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewUser new user
//
// swagger:model NewUser
type NewUser struct {

	// email
	// Example: john@email.com
	Email string `json:"email,omitempty"`

	// first name
	// Example: John
	FirstName string `json:"first_name,omitempty"`

	// last name
	// Example: Doe
	LastName string `json:"last_name,omitempty"`

	// login
	// Example: login
	Login string `json:"login,omitempty"`

	// password
	// Example: 12345
	Password string `json:"password,omitempty"`

	// telephone
	// Example: 12345
	Telephone string `json:"telephone,omitempty"`
}

// Validate validates this new user
func (m *NewUser) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this new user based on context it is used
func (m *NewUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewUser) UnmarshalBinary(b []byte) error {
	var res NewUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
