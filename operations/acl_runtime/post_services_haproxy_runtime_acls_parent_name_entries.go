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

package acl_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostServicesHaproxyRuntimeAclsParentNameEntriesHandlerFunc turns a function with the right signature into a post services haproxy runtime acls parent name entries handler
type PostServicesHaproxyRuntimeAclsParentNameEntriesHandlerFunc func(PostServicesHaproxyRuntimeAclsParentNameEntriesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostServicesHaproxyRuntimeAclsParentNameEntriesHandlerFunc) Handle(params PostServicesHaproxyRuntimeAclsParentNameEntriesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostServicesHaproxyRuntimeAclsParentNameEntriesHandler interface for that can handle valid post services haproxy runtime acls parent name entries params
type PostServicesHaproxyRuntimeAclsParentNameEntriesHandler interface {
	Handle(PostServicesHaproxyRuntimeAclsParentNameEntriesParams, interface{}) middleware.Responder
}

// NewPostServicesHaproxyRuntimeAclsParentNameEntries creates a new http.Handler for the post services haproxy runtime acls parent name entries operation
func NewPostServicesHaproxyRuntimeAclsParentNameEntries(ctx *middleware.Context, handler PostServicesHaproxyRuntimeAclsParentNameEntriesHandler) *PostServicesHaproxyRuntimeAclsParentNameEntries {
	return &PostServicesHaproxyRuntimeAclsParentNameEntries{Context: ctx, Handler: handler}
}

/*
	PostServicesHaproxyRuntimeAclsParentNameEntries swagger:route POST /services/haproxy/runtime/acls/{parent_name}/entries ACL Runtime postServicesHaproxyRuntimeAclsParentNameEntries

# Add entry to an ACL file

Adds an entry into the ACL file using the runtime socket.
*/
type PostServicesHaproxyRuntimeAclsParentNameEntries struct {
	Context *middleware.Context
	Handler PostServicesHaproxyRuntimeAclsParentNameEntriesHandler
}

func (o *PostServicesHaproxyRuntimeAclsParentNameEntries) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostServicesHaproxyRuntimeAclsParentNameEntriesParams()
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