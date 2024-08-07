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
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// CreateServerBackendCreatedCode is the HTTP code returned for type CreateServerBackendCreated
const CreateServerBackendCreatedCode int = 201

/*
CreateServerBackendCreated Server created

swagger:response createServerBackendCreated
*/
type CreateServerBackendCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Server `json:"body,omitempty"`
}

// NewCreateServerBackendCreated creates CreateServerBackendCreated with default headers values
func NewCreateServerBackendCreated() *CreateServerBackendCreated {

	return &CreateServerBackendCreated{}
}

// WithPayload adds the payload to the create server backend created response
func (o *CreateServerBackendCreated) WithPayload(payload *models.Server) *CreateServerBackendCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create server backend created response
func (o *CreateServerBackendCreated) SetPayload(payload *models.Server) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServerBackendCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateServerBackendAcceptedCode is the HTTP code returned for type CreateServerBackendAccepted
const CreateServerBackendAcceptedCode int = 202

/*
CreateServerBackendAccepted Configuration change accepted and reload requested

swagger:response createServerBackendAccepted
*/
type CreateServerBackendAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Server `json:"body,omitempty"`
}

// NewCreateServerBackendAccepted creates CreateServerBackendAccepted with default headers values
func NewCreateServerBackendAccepted() *CreateServerBackendAccepted {

	return &CreateServerBackendAccepted{}
}

// WithReloadID adds the reloadId to the create server backend accepted response
func (o *CreateServerBackendAccepted) WithReloadID(reloadID string) *CreateServerBackendAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create server backend accepted response
func (o *CreateServerBackendAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create server backend accepted response
func (o *CreateServerBackendAccepted) WithPayload(payload *models.Server) *CreateServerBackendAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create server backend accepted response
func (o *CreateServerBackendAccepted) SetPayload(payload *models.Server) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServerBackendAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateServerBackendBadRequestCode is the HTTP code returned for type CreateServerBackendBadRequest
const CreateServerBackendBadRequestCode int = 400

/*
CreateServerBackendBadRequest Bad request

swagger:response createServerBackendBadRequest
*/
type CreateServerBackendBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateServerBackendBadRequest creates CreateServerBackendBadRequest with default headers values
func NewCreateServerBackendBadRequest() *CreateServerBackendBadRequest {

	return &CreateServerBackendBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create server backend bad request response
func (o *CreateServerBackendBadRequest) WithConfigurationVersion(configurationVersion string) *CreateServerBackendBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create server backend bad request response
func (o *CreateServerBackendBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create server backend bad request response
func (o *CreateServerBackendBadRequest) WithPayload(payload *models.Error) *CreateServerBackendBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create server backend bad request response
func (o *CreateServerBackendBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServerBackendBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateServerBackendConflictCode is the HTTP code returned for type CreateServerBackendConflict
const CreateServerBackendConflictCode int = 409

/*
CreateServerBackendConflict The specified resource already exists

swagger:response createServerBackendConflict
*/
type CreateServerBackendConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateServerBackendConflict creates CreateServerBackendConflict with default headers values
func NewCreateServerBackendConflict() *CreateServerBackendConflict {

	return &CreateServerBackendConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create server backend conflict response
func (o *CreateServerBackendConflict) WithConfigurationVersion(configurationVersion string) *CreateServerBackendConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create server backend conflict response
func (o *CreateServerBackendConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create server backend conflict response
func (o *CreateServerBackendConflict) WithPayload(payload *models.Error) *CreateServerBackendConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create server backend conflict response
func (o *CreateServerBackendConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServerBackendConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
CreateServerBackendDefault General Error

swagger:response createServerBackendDefault
*/
type CreateServerBackendDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateServerBackendDefault creates CreateServerBackendDefault with default headers values
func NewCreateServerBackendDefault(code int) *CreateServerBackendDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateServerBackendDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create server backend default response
func (o *CreateServerBackendDefault) WithStatusCode(code int) *CreateServerBackendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create server backend default response
func (o *CreateServerBackendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create server backend default response
func (o *CreateServerBackendDefault) WithConfigurationVersion(configurationVersion string) *CreateServerBackendDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create server backend default response
func (o *CreateServerBackendDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create server backend default response
func (o *CreateServerBackendDefault) WithPayload(payload *models.Error) *CreateServerBackendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create server backend default response
func (o *CreateServerBackendDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServerBackendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
