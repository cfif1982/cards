// Code generated by go-swagger; DO NOT EDIT.

package card

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetCardByNumberHandlerFunc turns a function with the right signature into a get card by number handler
type GetCardByNumberHandlerFunc func(GetCardByNumberParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCardByNumberHandlerFunc) Handle(params GetCardByNumberParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetCardByNumberHandler interface for that can handle valid get card by number params
type GetCardByNumberHandler interface {
	Handle(GetCardByNumberParams, interface{}) middleware.Responder
}

// NewGetCardByNumber creates a new http.Handler for the get card by number operation
func NewGetCardByNumber(ctx *middleware.Context, handler GetCardByNumberHandler) *GetCardByNumber {
	return &GetCardByNumber{Context: ctx, Handler: handler}
}

/*
	GetCardByNumber swagger:route GET /card/{cardNumber} card getCardByNumber

# Find card by number

Returns a single card
*/
type GetCardByNumber struct {
	Context *middleware.Context
	Handler GetCardByNumberHandler
}

func (o *GetCardByNumber) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetCardByNumberParams()
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
