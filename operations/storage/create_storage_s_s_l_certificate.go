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

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateStorageSSLCertificateHandlerFunc turns a function with the right signature into a create storage s s l certificate handler
type CreateStorageSSLCertificateHandlerFunc func(CreateStorageSSLCertificateParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateStorageSSLCertificateHandlerFunc) Handle(params CreateStorageSSLCertificateParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateStorageSSLCertificateHandler interface for that can handle valid create storage s s l certificate params
type CreateStorageSSLCertificateHandler interface {
	Handle(CreateStorageSSLCertificateParams, interface{}) middleware.Responder
}

// NewCreateStorageSSLCertificate creates a new http.Handler for the create storage s s l certificate operation
func NewCreateStorageSSLCertificate(ctx *middleware.Context, handler CreateStorageSSLCertificateHandler) *CreateStorageSSLCertificate {
	return &CreateStorageSSLCertificate{Context: ctx, Handler: handler}
}

/*CreateStorageSSLCertificate swagger:route POST /services/haproxy/storage/ssl_certificates Storage createStorageSSLCertificate

Create SSL certificate

Creates SSL certificate.

*/
type CreateStorageSSLCertificate struct {
	Context *middleware.Context
	Handler CreateStorageSSLCertificateHandler
}

func (o *CreateStorageSSLCertificate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateStorageSSLCertificateParams()

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
