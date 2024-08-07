// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package tcp_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAllTCPResponseRuleBackendHandlerFunc turns a function with the right signature into a get all TCP response rule backend handler
type GetAllTCPResponseRuleBackendHandlerFunc func(GetAllTCPResponseRuleBackendParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllTCPResponseRuleBackendHandlerFunc) Handle(params GetAllTCPResponseRuleBackendParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetAllTCPResponseRuleBackendHandler interface for that can handle valid get all TCP response rule backend params
type GetAllTCPResponseRuleBackendHandler interface {
	Handle(GetAllTCPResponseRuleBackendParams, interface{}) middleware.Responder
}

// NewGetAllTCPResponseRuleBackend creates a new http.Handler for the get all TCP response rule backend operation
func NewGetAllTCPResponseRuleBackend(ctx *middleware.Context, handler GetAllTCPResponseRuleBackendHandler) *GetAllTCPResponseRuleBackend {
	return &GetAllTCPResponseRuleBackend{Context: ctx, Handler: handler}
}

/*
	GetAllTCPResponseRuleBackend swagger:route GET /services/haproxy/configuration/backends/{parent_name}/tcp_response_rules TCPResponseRule getAllTcpResponseRuleBackend

# Return an array of all TCP Response Rules

Returns all TCP Response Rules that are configured in specified backend.
*/
type GetAllTCPResponseRuleBackend struct {
	Context *middleware.Context
	Handler GetAllTCPResponseRuleBackendHandler
}

func (o *GetAllTCPResponseRuleBackend) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAllTCPResponseRuleBackendParams()
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
