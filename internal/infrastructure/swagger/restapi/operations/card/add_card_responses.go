// Code generated by go-swagger; DO NOT EDIT.

package card

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"cards/internal/infrastructure/swagger/models"
)

// AddCardOKCode is the HTTP code returned for type AddCardOK
const AddCardOKCode int = 200

/*
AddCardOK Successful operation

swagger:response addCardOK
*/
type AddCardOK struct {

	/*
	  In: Body
	*/
	Payload *models.Card `json:"body,omitempty"`
}

// NewAddCardOK creates AddCardOK with default headers values
func NewAddCardOK() *AddCardOK {

	return &AddCardOK{}
}

// WithPayload adds the payload to the add card o k response
func (o *AddCardOK) WithPayload(payload *models.Card) *AddCardOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add card o k response
func (o *AddCardOK) SetPayload(payload *models.Card) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddCardOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddCardBadRequestCode is the HTTP code returned for type AddCardBadRequest
const AddCardBadRequestCode int = 400

/*
AddCardBadRequest Invalid input

swagger:response addCardBadRequest
*/
type AddCardBadRequest struct {
}

// NewAddCardBadRequest creates AddCardBadRequest with default headers values
func NewAddCardBadRequest() *AddCardBadRequest {

	return &AddCardBadRequest{}
}

// WriteResponse to the client
func (o *AddCardBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// AddCardUnprocessableEntityCode is the HTTP code returned for type AddCardUnprocessableEntity
const AddCardUnprocessableEntityCode int = 422

/*
AddCardUnprocessableEntity Validation exception

swagger:response addCardUnprocessableEntity
*/
type AddCardUnprocessableEntity struct {
}

// NewAddCardUnprocessableEntity creates AddCardUnprocessableEntity with default headers values
func NewAddCardUnprocessableEntity() *AddCardUnprocessableEntity {

	return &AddCardUnprocessableEntity{}
}

// WriteResponse to the client
func (o *AddCardUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(422)
}