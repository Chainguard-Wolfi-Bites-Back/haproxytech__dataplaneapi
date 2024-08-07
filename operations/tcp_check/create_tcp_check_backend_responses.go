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

package tcp_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// CreateTCPCheckBackendCreatedCode is the HTTP code returned for type CreateTCPCheckBackendCreated
const CreateTCPCheckBackendCreatedCode int = 201

/*
CreateTCPCheckBackendCreated TCP check created

swagger:response createTcpCheckBackendCreated
*/
type CreateTCPCheckBackendCreated struct {

	/*
	  In: Body
	*/
	Payload *models.TCPCheck `json:"body,omitempty"`
}

// NewCreateTCPCheckBackendCreated creates CreateTCPCheckBackendCreated with default headers values
func NewCreateTCPCheckBackendCreated() *CreateTCPCheckBackendCreated {

	return &CreateTCPCheckBackendCreated{}
}

// WithPayload adds the payload to the create Tcp check backend created response
func (o *CreateTCPCheckBackendCreated) WithPayload(payload *models.TCPCheck) *CreateTCPCheckBackendCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Tcp check backend created response
func (o *CreateTCPCheckBackendCreated) SetPayload(payload *models.TCPCheck) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTCPCheckBackendCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateTCPCheckBackendAcceptedCode is the HTTP code returned for type CreateTCPCheckBackendAccepted
const CreateTCPCheckBackendAcceptedCode int = 202

/*
CreateTCPCheckBackendAccepted Configuration change accepted and reload requested

swagger:response createTcpCheckBackendAccepted
*/
type CreateTCPCheckBackendAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.TCPCheck `json:"body,omitempty"`
}

// NewCreateTCPCheckBackendAccepted creates CreateTCPCheckBackendAccepted with default headers values
func NewCreateTCPCheckBackendAccepted() *CreateTCPCheckBackendAccepted {

	return &CreateTCPCheckBackendAccepted{}
}

// WithReloadID adds the reloadId to the create Tcp check backend accepted response
func (o *CreateTCPCheckBackendAccepted) WithReloadID(reloadID string) *CreateTCPCheckBackendAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create Tcp check backend accepted response
func (o *CreateTCPCheckBackendAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create Tcp check backend accepted response
func (o *CreateTCPCheckBackendAccepted) WithPayload(payload *models.TCPCheck) *CreateTCPCheckBackendAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Tcp check backend accepted response
func (o *CreateTCPCheckBackendAccepted) SetPayload(payload *models.TCPCheck) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTCPCheckBackendAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateTCPCheckBackendBadRequestCode is the HTTP code returned for type CreateTCPCheckBackendBadRequest
const CreateTCPCheckBackendBadRequestCode int = 400

/*
CreateTCPCheckBackendBadRequest Bad request

swagger:response createTcpCheckBackendBadRequest
*/
type CreateTCPCheckBackendBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateTCPCheckBackendBadRequest creates CreateTCPCheckBackendBadRequest with default headers values
func NewCreateTCPCheckBackendBadRequest() *CreateTCPCheckBackendBadRequest {

	return &CreateTCPCheckBackendBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create Tcp check backend bad request response
func (o *CreateTCPCheckBackendBadRequest) WithConfigurationVersion(configurationVersion string) *CreateTCPCheckBackendBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create Tcp check backend bad request response
func (o *CreateTCPCheckBackendBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create Tcp check backend bad request response
func (o *CreateTCPCheckBackendBadRequest) WithPayload(payload *models.Error) *CreateTCPCheckBackendBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Tcp check backend bad request response
func (o *CreateTCPCheckBackendBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTCPCheckBackendBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateTCPCheckBackendConflictCode is the HTTP code returned for type CreateTCPCheckBackendConflict
const CreateTCPCheckBackendConflictCode int = 409

/*
CreateTCPCheckBackendConflict The specified resource already exists

swagger:response createTcpCheckBackendConflict
*/
type CreateTCPCheckBackendConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateTCPCheckBackendConflict creates CreateTCPCheckBackendConflict with default headers values
func NewCreateTCPCheckBackendConflict() *CreateTCPCheckBackendConflict {

	return &CreateTCPCheckBackendConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create Tcp check backend conflict response
func (o *CreateTCPCheckBackendConflict) WithConfigurationVersion(configurationVersion string) *CreateTCPCheckBackendConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create Tcp check backend conflict response
func (o *CreateTCPCheckBackendConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create Tcp check backend conflict response
func (o *CreateTCPCheckBackendConflict) WithPayload(payload *models.Error) *CreateTCPCheckBackendConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Tcp check backend conflict response
func (o *CreateTCPCheckBackendConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTCPCheckBackendConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
CreateTCPCheckBackendDefault General Error

swagger:response createTcpCheckBackendDefault
*/
type CreateTCPCheckBackendDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateTCPCheckBackendDefault creates CreateTCPCheckBackendDefault with default headers values
func NewCreateTCPCheckBackendDefault(code int) *CreateTCPCheckBackendDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateTCPCheckBackendDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create TCP check backend default response
func (o *CreateTCPCheckBackendDefault) WithStatusCode(code int) *CreateTCPCheckBackendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create TCP check backend default response
func (o *CreateTCPCheckBackendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create TCP check backend default response
func (o *CreateTCPCheckBackendDefault) WithConfigurationVersion(configurationVersion string) *CreateTCPCheckBackendDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create TCP check backend default response
func (o *CreateTCPCheckBackendDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create TCP check backend default response
func (o *CreateTCPCheckBackendDefault) WithPayload(payload *models.Error) *CreateTCPCheckBackendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create TCP check backend default response
func (o *CreateTCPCheckBackendDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTCPCheckBackendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
