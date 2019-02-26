// Code generated by go-swagger; DO NOT EDIT.

package log_target

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateLogTargetHandlerFunc turns a function with the right signature into a create log target handler
type CreateLogTargetHandlerFunc func(CreateLogTargetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateLogTargetHandlerFunc) Handle(params CreateLogTargetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateLogTargetHandler interface for that can handle valid create log target params
type CreateLogTargetHandler interface {
	Handle(CreateLogTargetParams, interface{}) middleware.Responder
}

// NewCreateLogTarget creates a new http.Handler for the create log target operation
func NewCreateLogTarget(ctx *middleware.Context, handler CreateLogTargetHandler) *CreateLogTarget {
	return &CreateLogTarget{Context: ctx, Handler: handler}
}

/*CreateLogTarget swagger:route POST /services/haproxy/configuration/log_targets LogTarget createLogTarget

Add a new Log Target

Adds a new Log Target of the specified type in the specified parent.

*/
type CreateLogTarget struct {
	Context *middleware.Context
	Handler CreateLogTargetHandler
}

func (o *CreateLogTarget) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateLogTargetParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}