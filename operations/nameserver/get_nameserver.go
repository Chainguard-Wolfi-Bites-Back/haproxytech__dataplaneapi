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

package nameserver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models/v2"
)

// GetNameserverHandlerFunc turns a function with the right signature into a get nameserver handler
type GetNameserverHandlerFunc func(GetNameserverParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNameserverHandlerFunc) Handle(params GetNameserverParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetNameserverHandler interface for that can handle valid get nameserver params
type GetNameserverHandler interface {
	Handle(GetNameserverParams, interface{}) middleware.Responder
}

// NewGetNameserver creates a new http.Handler for the get nameserver operation
func NewGetNameserver(ctx *middleware.Context, handler GetNameserverHandler) *GetNameserver {
	return &GetNameserver{Context: ctx, Handler: handler}
}

/*GetNameserver swagger:route GET /services/haproxy/configuration/nameservers/{name} Nameserver getNameserver

Return a nameserver

Returns one nameserver configuration by it's name.

*/
type GetNameserver struct {
	Context *middleware.Context
	Handler GetNameserverHandler
}

func (o *GetNameserver) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetNameserverParams()

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

// GetNameserverOKBody get nameserver o k body
//
// swagger:model GetNameserverOKBody
type GetNameserverOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.Nameserver `json:"data,omitempty"`
}

// Validate validates this get nameserver o k body
func (o *GetNameserverOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNameserverOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNameserverOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNameserverOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNameserverOKBody) UnmarshalBinary(b []byte) error {
	var res GetNameserverOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
