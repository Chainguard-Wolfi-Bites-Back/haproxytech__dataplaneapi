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

	"github.com/haproxytech/client-native/v2/models"
)

// CreateTCPCheckCreatedCode is the HTTP code returned for type CreateTCPCheckCreated
const CreateTCPCheckCreatedCode int = 201

/*CreateTCPCheckCreated TCP check created

swagger:response createTcpCheckCreated
*/
type CreateTCPCheckCreated struct {

	/*
	  In: Body
	*/
	Payload *models.TCPCheck `json:"body,omitempty"`
}

// NewCreateTCPCheckCreated creates CreateTCPCheckCreated with default headers values
func NewCreateTCPCheckCreated() *CreateTCPCheckCreated {

	return &CreateTCPCheckCreated{}
}

// WithPayload adds the payload to the create Tcp check created response
func (o *CreateTCPCheckCreated) WithPayload(payload *models.TCPCheck) *CreateTCPCheckCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Tcp check created response
func (o *CreateTCPCheckCreated) SetPayload(payload *models.TCPCheck) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTCPCheckCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateTCPCheckAcceptedCode is the HTTP code returned for type CreateTCPCheckAccepted
const CreateTCPCheckAcceptedCode int = 202

/*CreateTCPCheckAccepted Configuration change accepted and reload requested

swagger:response createTcpCheckAccepted
*/
type CreateTCPCheckAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.TCPCheck `json:"body,omitempty"`
}

// NewCreateTCPCheckAccepted creates CreateTCPCheckAccepted with default headers values
func NewCreateTCPCheckAccepted() *CreateTCPCheckAccepted {

	return &CreateTCPCheckAccepted{}
}

// WithReloadID adds the reloadId to the create Tcp check accepted response
func (o *CreateTCPCheckAccepted) WithReloadID(reloadID string) *CreateTCPCheckAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create Tcp check accepted response
func (o *CreateTCPCheckAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create Tcp check accepted response
func (o *CreateTCPCheckAccepted) WithPayload(payload *models.TCPCheck) *CreateTCPCheckAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Tcp check accepted response
func (o *CreateTCPCheckAccepted) SetPayload(payload *models.TCPCheck) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTCPCheckAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateTCPCheckBadRequestCode is the HTTP code returned for type CreateTCPCheckBadRequest
const CreateTCPCheckBadRequestCode int = 400

/*CreateTCPCheckBadRequest Bad request

swagger:response createTcpCheckBadRequest
*/
type CreateTCPCheckBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateTCPCheckBadRequest creates CreateTCPCheckBadRequest with default headers values
func NewCreateTCPCheckBadRequest() *CreateTCPCheckBadRequest {

	return &CreateTCPCheckBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create Tcp check bad request response
func (o *CreateTCPCheckBadRequest) WithConfigurationVersion(configurationVersion string) *CreateTCPCheckBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create Tcp check bad request response
func (o *CreateTCPCheckBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create Tcp check bad request response
func (o *CreateTCPCheckBadRequest) WithPayload(payload *models.Error) *CreateTCPCheckBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Tcp check bad request response
func (o *CreateTCPCheckBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTCPCheckBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateTCPCheckConflictCode is the HTTP code returned for type CreateTCPCheckConflict
const CreateTCPCheckConflictCode int = 409

/*CreateTCPCheckConflict The specified resource already exists

swagger:response createTcpCheckConflict
*/
type CreateTCPCheckConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateTCPCheckConflict creates CreateTCPCheckConflict with default headers values
func NewCreateTCPCheckConflict() *CreateTCPCheckConflict {

	return &CreateTCPCheckConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create Tcp check conflict response
func (o *CreateTCPCheckConflict) WithConfigurationVersion(configurationVersion string) *CreateTCPCheckConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create Tcp check conflict response
func (o *CreateTCPCheckConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create Tcp check conflict response
func (o *CreateTCPCheckConflict) WithPayload(payload *models.Error) *CreateTCPCheckConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Tcp check conflict response
func (o *CreateTCPCheckConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTCPCheckConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*CreateTCPCheckDefault General Error

swagger:response createTcpCheckDefault
*/
type CreateTCPCheckDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateTCPCheckDefault creates CreateTCPCheckDefault with default headers values
func NewCreateTCPCheckDefault(code int) *CreateTCPCheckDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateTCPCheckDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create TCP check default response
func (o *CreateTCPCheckDefault) WithStatusCode(code int) *CreateTCPCheckDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create TCP check default response
func (o *CreateTCPCheckDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create TCP check default response
func (o *CreateTCPCheckDefault) WithConfigurationVersion(configurationVersion string) *CreateTCPCheckDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create TCP check default response
func (o *CreateTCPCheckDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create TCP check default response
func (o *CreateTCPCheckDefault) WithPayload(payload *models.Error) *CreateTCPCheckDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create TCP check default response
func (o *CreateTCPCheckDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTCPCheckDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
