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

// ReplaceTCPResponseRuleBackendHandlerFunc turns a function with the right signature into a replace TCP response rule backend handler
type ReplaceTCPResponseRuleBackendHandlerFunc func(ReplaceTCPResponseRuleBackendParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceTCPResponseRuleBackendHandlerFunc) Handle(params ReplaceTCPResponseRuleBackendParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceTCPResponseRuleBackendHandler interface for that can handle valid replace TCP response rule backend params
type ReplaceTCPResponseRuleBackendHandler interface {
	Handle(ReplaceTCPResponseRuleBackendParams, interface{}) middleware.Responder
}

// NewReplaceTCPResponseRuleBackend creates a new http.Handler for the replace TCP response rule backend operation
func NewReplaceTCPResponseRuleBackend(ctx *middleware.Context, handler ReplaceTCPResponseRuleBackendHandler) *ReplaceTCPResponseRuleBackend {
	return &ReplaceTCPResponseRuleBackend{Context: ctx, Handler: handler}
}

/*
	ReplaceTCPResponseRuleBackend swagger:route PUT /services/haproxy/configuration/backends/{parent_name}/tcp_response_rules/{index} TCPResponseRule replaceTcpResponseRuleBackend

# Replace a TCP Response Rule

Replaces a TCP Response Rule configuration by it's Index in the specified backend.
*/
type ReplaceTCPResponseRuleBackend struct {
	Context *middleware.Context
	Handler ReplaceTCPResponseRuleBackendHandler
}

func (o *ReplaceTCPResponseRuleBackend) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewReplaceTCPResponseRuleBackendParams()
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