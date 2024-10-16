// Code generated by go-swagger; DO NOT EDIT.

package bank

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateBankHandlerFunc turns a function with the right signature into a update bank handler
type UpdateBankHandlerFunc func(UpdateBankParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateBankHandlerFunc) Handle(params UpdateBankParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateBankHandler interface for that can handle valid update bank params
type UpdateBankHandler interface {
	Handle(UpdateBankParams, interface{}) middleware.Responder
}

// NewUpdateBank creates a new http.Handler for the update bank operation
func NewUpdateBank(ctx *middleware.Context, handler UpdateBankHandler) *UpdateBank {
	return &UpdateBank{Context: ctx, Handler: handler}
}

/*
	UpdateBank swagger:route PUT /bank bank updateBank

# Update an existing bank

Update an existing bank by uuid
*/
type UpdateBank struct {
	Context *middleware.Context
	Handler UpdateBankHandler
}

func (o *UpdateBank) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateBankParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
