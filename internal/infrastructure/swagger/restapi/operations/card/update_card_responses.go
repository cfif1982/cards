// Code generated by go-swagger; DO NOT EDIT.

package card

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"cards/internal/infrastructure/swagger/models"
)

// UpdateCardOKCode is the HTTP code returned for type UpdateCardOK
const UpdateCardOKCode int = 200

/*
UpdateCardOK Successful operation

swagger:response updateCardOK
*/
type UpdateCardOK struct {

	/*
	  In: Body
	*/
	Payload *models.Card `json:"body,omitempty"`
}

// NewUpdateCardOK creates UpdateCardOK with default headers values
func NewUpdateCardOK() *UpdateCardOK {

	return &UpdateCardOK{}
}

// WithPayload adds the payload to the update card o k response
func (o *UpdateCardOK) WithPayload(payload *models.Card) *UpdateCardOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update card o k response
func (o *UpdateCardOK) SetPayload(payload *models.Card) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCardOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCardBadRequestCode is the HTTP code returned for type UpdateCardBadRequest
const UpdateCardBadRequestCode int = 400

/*
UpdateCardBadRequest Invalid ID supplied

swagger:response updateCardBadRequest
*/
type UpdateCardBadRequest struct {
}

// NewUpdateCardBadRequest creates UpdateCardBadRequest with default headers values
func NewUpdateCardBadRequest() *UpdateCardBadRequest {

	return &UpdateCardBadRequest{}
}

// WriteResponse to the client
func (o *UpdateCardBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UpdateCardNotFoundCode is the HTTP code returned for type UpdateCardNotFound
const UpdateCardNotFoundCode int = 404

/*
UpdateCardNotFound Card not found

swagger:response updateCardNotFound
*/
type UpdateCardNotFound struct {
}

// NewUpdateCardNotFound creates UpdateCardNotFound with default headers values
func NewUpdateCardNotFound() *UpdateCardNotFound {

	return &UpdateCardNotFound{}
}

// WriteResponse to the client
func (o *UpdateCardNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// UpdateCardUnprocessableEntityCode is the HTTP code returned for type UpdateCardUnprocessableEntity
const UpdateCardUnprocessableEntityCode int = 422

/*
UpdateCardUnprocessableEntity Validation exception

swagger:response updateCardUnprocessableEntity
*/
type UpdateCardUnprocessableEntity struct {
}

// NewUpdateCardUnprocessableEntity creates UpdateCardUnprocessableEntity with default headers values
func NewUpdateCardUnprocessableEntity() *UpdateCardUnprocessableEntity {

	return &UpdateCardUnprocessableEntity{}
}

// WriteResponse to the client
func (o *UpdateCardUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(422)
}
