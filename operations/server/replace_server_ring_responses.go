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

// ReplaceServerRingOKCode is the HTTP code returned for type ReplaceServerRingOK
const ReplaceServerRingOKCode int = 200

/*
ReplaceServerRingOK Server replaced

swagger:response replaceServerRingOK
*/
type ReplaceServerRingOK struct {

	/*
	  In: Body
	*/
	Payload *models.Server `json:"body,omitempty"`
}

// NewReplaceServerRingOK creates ReplaceServerRingOK with default headers values
func NewReplaceServerRingOK() *ReplaceServerRingOK {

	return &ReplaceServerRingOK{}
}

// WithPayload adds the payload to the replace server ring o k response
func (o *ReplaceServerRingOK) WithPayload(payload *models.Server) *ReplaceServerRingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server ring o k response
func (o *ReplaceServerRingOK) SetPayload(payload *models.Server) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerRingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceServerRingAcceptedCode is the HTTP code returned for type ReplaceServerRingAccepted
const ReplaceServerRingAcceptedCode int = 202

/*
ReplaceServerRingAccepted Configuration change accepted and reload requested

swagger:response replaceServerRingAccepted
*/
type ReplaceServerRingAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Server `json:"body,omitempty"`
}

// NewReplaceServerRingAccepted creates ReplaceServerRingAccepted with default headers values
func NewReplaceServerRingAccepted() *ReplaceServerRingAccepted {

	return &ReplaceServerRingAccepted{}
}

// WithReloadID adds the reloadId to the replace server ring accepted response
func (o *ReplaceServerRingAccepted) WithReloadID(reloadID string) *ReplaceServerRingAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace server ring accepted response
func (o *ReplaceServerRingAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace server ring accepted response
func (o *ReplaceServerRingAccepted) WithPayload(payload *models.Server) *ReplaceServerRingAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server ring accepted response
func (o *ReplaceServerRingAccepted) SetPayload(payload *models.Server) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerRingAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceServerRingBadRequestCode is the HTTP code returned for type ReplaceServerRingBadRequest
const ReplaceServerRingBadRequestCode int = 400

/*
ReplaceServerRingBadRequest Bad request

swagger:response replaceServerRingBadRequest
*/
type ReplaceServerRingBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceServerRingBadRequest creates ReplaceServerRingBadRequest with default headers values
func NewReplaceServerRingBadRequest() *ReplaceServerRingBadRequest {

	return &ReplaceServerRingBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace server ring bad request response
func (o *ReplaceServerRingBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceServerRingBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace server ring bad request response
func (o *ReplaceServerRingBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace server ring bad request response
func (o *ReplaceServerRingBadRequest) WithPayload(payload *models.Error) *ReplaceServerRingBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server ring bad request response
func (o *ReplaceServerRingBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerRingBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceServerRingNotFoundCode is the HTTP code returned for type ReplaceServerRingNotFound
const ReplaceServerRingNotFoundCode int = 404

/*
ReplaceServerRingNotFound The specified resource was not found

swagger:response replaceServerRingNotFound
*/
type ReplaceServerRingNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceServerRingNotFound creates ReplaceServerRingNotFound with default headers values
func NewReplaceServerRingNotFound() *ReplaceServerRingNotFound {

	return &ReplaceServerRingNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the replace server ring not found response
func (o *ReplaceServerRingNotFound) WithConfigurationVersion(configurationVersion string) *ReplaceServerRingNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace server ring not found response
func (o *ReplaceServerRingNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace server ring not found response
func (o *ReplaceServerRingNotFound) WithPayload(payload *models.Error) *ReplaceServerRingNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server ring not found response
func (o *ReplaceServerRingNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerRingNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
ReplaceServerRingDefault General Error

swagger:response replaceServerRingDefault
*/
type ReplaceServerRingDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceServerRingDefault creates ReplaceServerRingDefault with default headers values
func NewReplaceServerRingDefault(code int) *ReplaceServerRingDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceServerRingDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace server ring default response
func (o *ReplaceServerRingDefault) WithStatusCode(code int) *ReplaceServerRingDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace server ring default response
func (o *ReplaceServerRingDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace server ring default response
func (o *ReplaceServerRingDefault) WithConfigurationVersion(configurationVersion string) *ReplaceServerRingDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace server ring default response
func (o *ReplaceServerRingDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace server ring default response
func (o *ReplaceServerRingDefault) WithPayload(payload *models.Error) *ReplaceServerRingDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server ring default response
func (o *ReplaceServerRingDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerRingDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
