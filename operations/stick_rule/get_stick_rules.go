// Code generated by go-swagger; DO NOT EDIT.

package stick_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetStickRulesHandlerFunc turns a function with the right signature into a get stick rules handler
type GetStickRulesHandlerFunc func(GetStickRulesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetStickRulesHandlerFunc) Handle(params GetStickRulesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetStickRulesHandler interface for that can handle valid get stick rules params
type GetStickRulesHandler interface {
	Handle(GetStickRulesParams, interface{}) middleware.Responder
}

// NewGetStickRules creates a new http.Handler for the get stick rules operation
func NewGetStickRules(ctx *middleware.Context, handler GetStickRulesHandler) *GetStickRules {
	return &GetStickRules{Context: ctx, Handler: handler}
}

/*GetStickRules swagger:route GET /services/haproxy/configuration/stick_rules StickRule getStickRules

Return an array of all Stick Rules

Returns all Stick Rules that are configured in specified backend.

*/
type GetStickRules struct {
	Context *middleware.Context
	Handler GetStickRulesHandler
}

func (o *GetStickRules) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetStickRulesParams()

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