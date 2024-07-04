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

package log_target

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAllLogTargetGlobalHandlerFunc turns a function with the right signature into a get all log target global handler
type GetAllLogTargetGlobalHandlerFunc func(GetAllLogTargetGlobalParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllLogTargetGlobalHandlerFunc) Handle(params GetAllLogTargetGlobalParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetAllLogTargetGlobalHandler interface for that can handle valid get all log target global params
type GetAllLogTargetGlobalHandler interface {
	Handle(GetAllLogTargetGlobalParams, interface{}) middleware.Responder
}

// NewGetAllLogTargetGlobal creates a new http.Handler for the get all log target global operation
func NewGetAllLogTargetGlobal(ctx *middleware.Context, handler GetAllLogTargetGlobalHandler) *GetAllLogTargetGlobal {
	return &GetAllLogTargetGlobal{Context: ctx, Handler: handler}
}

/*
	GetAllLogTargetGlobal swagger:route GET /services/haproxy/configuration/global/log_targets LogTarget getAllLogTargetGlobal

# Return an array of all Log Targets

Returns all Log Targets that are configured in specified parent.
*/
type GetAllLogTargetGlobal struct {
	Context *middleware.Context
	Handler GetAllLogTargetGlobalHandler
}

func (o *GetAllLogTargetGlobal) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAllLogTargetGlobalParams()
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