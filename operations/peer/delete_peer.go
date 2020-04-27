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

package peer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeletePeerHandlerFunc turns a function with the right signature into a delete peer handler
type DeletePeerHandlerFunc func(DeletePeerParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeletePeerHandlerFunc) Handle(params DeletePeerParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeletePeerHandler interface for that can handle valid delete peer params
type DeletePeerHandler interface {
	Handle(DeletePeerParams, interface{}) middleware.Responder
}

// NewDeletePeer creates a new http.Handler for the delete peer operation
func NewDeletePeer(ctx *middleware.Context, handler DeletePeerHandler) *DeletePeer {
	return &DeletePeer{Context: ctx, Handler: handler}
}

/*DeletePeer swagger:route DELETE /services/haproxy/configuration/peer_section/{name} Peer deletePeer

Delete a peer

Deletes a peer from the configuration by it's name.

*/
type DeletePeer struct {
	Context *middleware.Context
	Handler DeletePeerHandler
}

func (o *DeletePeer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeletePeerParams()

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
