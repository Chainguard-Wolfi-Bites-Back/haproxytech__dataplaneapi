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

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/haproxytech/models/v2"
)

// GetServersHandlerFunc turns a function with the right signature into a get servers handler
type GetServersHandlerFunc func(GetServersParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetServersHandlerFunc) Handle(params GetServersParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetServersHandler interface for that can handle valid get servers params
type GetServersHandler interface {
	Handle(GetServersParams, interface{}) middleware.Responder
}

// NewGetServers creates a new http.Handler for the get servers operation
func NewGetServers(ctx *middleware.Context, handler GetServersHandler) *GetServers {
	return &GetServers{Context: ctx, Handler: handler}
}

/*GetServers swagger:route GET /services/haproxy/configuration/servers Server getServers

Return an array of servers

Returns an array of all servers that are configured in specified backend.

*/
type GetServers struct {
	Context *middleware.Context
	Handler GetServersHandler
}

func (o *GetServers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetServersParams()

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

// GetServersOKBody get servers o k body
//
// swagger:model GetServersOKBody
type GetServersOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.Servers `json:"data"`
}

// Validate validates this get servers o k body
func (o *GetServersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServersOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getServersOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getServersOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServersOKBody) UnmarshalBinary(b []byte) error {
	var res GetServersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
