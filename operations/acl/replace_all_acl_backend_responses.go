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

package acl

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// ReplaceAllACLBackendOKCode is the HTTP code returned for type ReplaceAllACLBackendOK
const ReplaceAllACLBackendOKCode int = 200

/*
ReplaceAllACLBackendOK All ACL lines replaced

swagger:response replaceAllAclBackendOK
*/
type ReplaceAllACLBackendOK struct {

	/*
	  In: Body
	*/
	Payload models.Acls `json:"body,omitempty"`
}

// NewReplaceAllACLBackendOK creates ReplaceAllACLBackendOK with default headers values
func NewReplaceAllACLBackendOK() *ReplaceAllACLBackendOK {

	return &ReplaceAllACLBackendOK{}
}

// WithPayload adds the payload to the replace all Acl backend o k response
func (o *ReplaceAllACLBackendOK) WithPayload(payload models.Acls) *ReplaceAllACLBackendOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace all Acl backend o k response
func (o *ReplaceAllACLBackendOK) SetPayload(payload models.Acls) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAllACLBackendOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Acls{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ReplaceAllACLBackendAcceptedCode is the HTTP code returned for type ReplaceAllACLBackendAccepted
const ReplaceAllACLBackendAcceptedCode int = 202

/*
ReplaceAllACLBackendAccepted Configuration change accepted and reload requested

swagger:response replaceAllAclBackendAccepted
*/
type ReplaceAllACLBackendAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload models.Acls `json:"body,omitempty"`
}

// NewReplaceAllACLBackendAccepted creates ReplaceAllACLBackendAccepted with default headers values
func NewReplaceAllACLBackendAccepted() *ReplaceAllACLBackendAccepted {

	return &ReplaceAllACLBackendAccepted{}
}

// WithReloadID adds the reloadId to the replace all Acl backend accepted response
func (o *ReplaceAllACLBackendAccepted) WithReloadID(reloadID string) *ReplaceAllACLBackendAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace all Acl backend accepted response
func (o *ReplaceAllACLBackendAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace all Acl backend accepted response
func (o *ReplaceAllACLBackendAccepted) WithPayload(payload models.Acls) *ReplaceAllACLBackendAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace all Acl backend accepted response
func (o *ReplaceAllACLBackendAccepted) SetPayload(payload models.Acls) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAllACLBackendAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Acls{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ReplaceAllACLBackendBadRequestCode is the HTTP code returned for type ReplaceAllACLBackendBadRequest
const ReplaceAllACLBackendBadRequestCode int = 400

/*
ReplaceAllACLBackendBadRequest Bad request

swagger:response replaceAllAclBackendBadRequest
*/
type ReplaceAllACLBackendBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceAllACLBackendBadRequest creates ReplaceAllACLBackendBadRequest with default headers values
func NewReplaceAllACLBackendBadRequest() *ReplaceAllACLBackendBadRequest {

	return &ReplaceAllACLBackendBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace all Acl backend bad request response
func (o *ReplaceAllACLBackendBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceAllACLBackendBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace all Acl backend bad request response
func (o *ReplaceAllACLBackendBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace all Acl backend bad request response
func (o *ReplaceAllACLBackendBadRequest) WithPayload(payload *models.Error) *ReplaceAllACLBackendBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace all Acl backend bad request response
func (o *ReplaceAllACLBackendBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAllACLBackendBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
ReplaceAllACLBackendDefault General Error

swagger:response replaceAllAclBackendDefault
*/
type ReplaceAllACLBackendDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceAllACLBackendDefault creates ReplaceAllACLBackendDefault with default headers values
func NewReplaceAllACLBackendDefault(code int) *ReplaceAllACLBackendDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceAllACLBackendDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace all Acl backend default response
func (o *ReplaceAllACLBackendDefault) WithStatusCode(code int) *ReplaceAllACLBackendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace all Acl backend default response
func (o *ReplaceAllACLBackendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace all Acl backend default response
func (o *ReplaceAllACLBackendDefault) WithConfigurationVersion(configurationVersion string) *ReplaceAllACLBackendDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace all Acl backend default response
func (o *ReplaceAllACLBackendDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace all Acl backend default response
func (o *ReplaceAllACLBackendDefault) WithPayload(payload *models.Error) *ReplaceAllACLBackendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace all Acl backend default response
func (o *ReplaceAllACLBackendDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAllACLBackendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}