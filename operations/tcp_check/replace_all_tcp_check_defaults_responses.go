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

// ReplaceAllTCPCheckDefaultsOKCode is the HTTP code returned for type ReplaceAllTCPCheckDefaultsOK
const ReplaceAllTCPCheckDefaultsOKCode int = 200

/*
ReplaceAllTCPCheckDefaultsOK All TCP Check lines replaced

swagger:response replaceAllTcpCheckDefaultsOK
*/
type ReplaceAllTCPCheckDefaultsOK struct {

	/*
	  In: Body
	*/
	Payload models.TCPChecks `json:"body,omitempty"`
}

// NewReplaceAllTCPCheckDefaultsOK creates ReplaceAllTCPCheckDefaultsOK with default headers values
func NewReplaceAllTCPCheckDefaultsOK() *ReplaceAllTCPCheckDefaultsOK {

	return &ReplaceAllTCPCheckDefaultsOK{}
}

// WithPayload adds the payload to the replace all Tcp check defaults o k response
func (o *ReplaceAllTCPCheckDefaultsOK) WithPayload(payload models.TCPChecks) *ReplaceAllTCPCheckDefaultsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace all Tcp check defaults o k response
func (o *ReplaceAllTCPCheckDefaultsOK) SetPayload(payload models.TCPChecks) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAllTCPCheckDefaultsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.TCPChecks{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ReplaceAllTCPCheckDefaultsAcceptedCode is the HTTP code returned for type ReplaceAllTCPCheckDefaultsAccepted
const ReplaceAllTCPCheckDefaultsAcceptedCode int = 202

/*
ReplaceAllTCPCheckDefaultsAccepted Configuration change accepted and reload requested

swagger:response replaceAllTcpCheckDefaultsAccepted
*/
type ReplaceAllTCPCheckDefaultsAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload models.TCPChecks `json:"body,omitempty"`
}

// NewReplaceAllTCPCheckDefaultsAccepted creates ReplaceAllTCPCheckDefaultsAccepted with default headers values
func NewReplaceAllTCPCheckDefaultsAccepted() *ReplaceAllTCPCheckDefaultsAccepted {

	return &ReplaceAllTCPCheckDefaultsAccepted{}
}

// WithReloadID adds the reloadId to the replace all Tcp check defaults accepted response
func (o *ReplaceAllTCPCheckDefaultsAccepted) WithReloadID(reloadID string) *ReplaceAllTCPCheckDefaultsAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace all Tcp check defaults accepted response
func (o *ReplaceAllTCPCheckDefaultsAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace all Tcp check defaults accepted response
func (o *ReplaceAllTCPCheckDefaultsAccepted) WithPayload(payload models.TCPChecks) *ReplaceAllTCPCheckDefaultsAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace all Tcp check defaults accepted response
func (o *ReplaceAllTCPCheckDefaultsAccepted) SetPayload(payload models.TCPChecks) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAllTCPCheckDefaultsAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.TCPChecks{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ReplaceAllTCPCheckDefaultsBadRequestCode is the HTTP code returned for type ReplaceAllTCPCheckDefaultsBadRequest
const ReplaceAllTCPCheckDefaultsBadRequestCode int = 400

/*
ReplaceAllTCPCheckDefaultsBadRequest Bad request

swagger:response replaceAllTcpCheckDefaultsBadRequest
*/
type ReplaceAllTCPCheckDefaultsBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceAllTCPCheckDefaultsBadRequest creates ReplaceAllTCPCheckDefaultsBadRequest with default headers values
func NewReplaceAllTCPCheckDefaultsBadRequest() *ReplaceAllTCPCheckDefaultsBadRequest {

	return &ReplaceAllTCPCheckDefaultsBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace all Tcp check defaults bad request response
func (o *ReplaceAllTCPCheckDefaultsBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceAllTCPCheckDefaultsBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace all Tcp check defaults bad request response
func (o *ReplaceAllTCPCheckDefaultsBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace all Tcp check defaults bad request response
func (o *ReplaceAllTCPCheckDefaultsBadRequest) WithPayload(payload *models.Error) *ReplaceAllTCPCheckDefaultsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace all Tcp check defaults bad request response
func (o *ReplaceAllTCPCheckDefaultsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAllTCPCheckDefaultsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
ReplaceAllTCPCheckDefaultsDefault General Error

swagger:response replaceAllTcpCheckDefaultsDefault
*/
type ReplaceAllTCPCheckDefaultsDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceAllTCPCheckDefaultsDefault creates ReplaceAllTCPCheckDefaultsDefault with default headers values
func NewReplaceAllTCPCheckDefaultsDefault(code int) *ReplaceAllTCPCheckDefaultsDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceAllTCPCheckDefaultsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace all TCP check defaults default response
func (o *ReplaceAllTCPCheckDefaultsDefault) WithStatusCode(code int) *ReplaceAllTCPCheckDefaultsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace all TCP check defaults default response
func (o *ReplaceAllTCPCheckDefaultsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace all TCP check defaults default response
func (o *ReplaceAllTCPCheckDefaultsDefault) WithConfigurationVersion(configurationVersion string) *ReplaceAllTCPCheckDefaultsDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace all TCP check defaults default response
func (o *ReplaceAllTCPCheckDefaultsDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace all TCP check defaults default response
func (o *ReplaceAllTCPCheckDefaultsDefault) WithPayload(payload *models.Error) *ReplaceAllTCPCheckDefaultsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace all TCP check defaults default response
func (o *ReplaceAllTCPCheckDefaultsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAllTCPCheckDefaultsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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