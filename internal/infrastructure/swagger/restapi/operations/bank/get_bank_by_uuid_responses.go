// Code generated by go-swagger; DO NOT EDIT.

package bank

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"cards/internal/infrastructure/swagger/models"
)

// GetBankByUUIDOKCode is the HTTP code returned for type GetBankByUUIDOK
const GetBankByUUIDOKCode int = 200

/*
GetBankByUUIDOK successful operation

swagger:response getBankByUuidOK
*/
type GetBankByUUIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Bank `json:"body,omitempty"`
}

// NewGetBankByUUIDOK creates GetBankByUUIDOK with default headers values
func NewGetBankByUUIDOK() *GetBankByUUIDOK {

	return &GetBankByUUIDOK{}
}

// WithPayload adds the payload to the get bank by Uuid o k response
func (o *GetBankByUUIDOK) WithPayload(payload *models.Bank) *GetBankByUUIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bank by Uuid o k response
func (o *GetBankByUUIDOK) SetPayload(payload *models.Bank) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBankByUUIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBankByUUIDBadRequestCode is the HTTP code returned for type GetBankByUUIDBadRequest
const GetBankByUUIDBadRequestCode int = 400

/*
GetBankByUUIDBadRequest Invalid ID supplied

swagger:response getBankByUuidBadRequest
*/
type GetBankByUUIDBadRequest struct {
}

// NewGetBankByUUIDBadRequest creates GetBankByUUIDBadRequest with default headers values
func NewGetBankByUUIDBadRequest() *GetBankByUUIDBadRequest {

	return &GetBankByUUIDBadRequest{}
}

// WriteResponse to the client
func (o *GetBankByUUIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetBankByUUIDNotFoundCode is the HTTP code returned for type GetBankByUUIDNotFound
const GetBankByUUIDNotFoundCode int = 404

/*
GetBankByUUIDNotFound Bank not found

swagger:response getBankByUuidNotFound
*/
type GetBankByUUIDNotFound struct {
}

// NewGetBankByUUIDNotFound creates GetBankByUUIDNotFound with default headers values
func NewGetBankByUUIDNotFound() *GetBankByUUIDNotFound {

	return &GetBankByUUIDNotFound{}
}

// WriteResponse to the client
func (o *GetBankByUUIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
