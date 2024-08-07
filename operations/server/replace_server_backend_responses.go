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

// ReplaceServerBackendOKCode is the HTTP code returned for type ReplaceServerBackendOK
const ReplaceServerBackendOKCode int = 200

/*
ReplaceServerBackendOK Server replaced

swagger:response replaceServerBackendOK
*/
type ReplaceServerBackendOK struct {

	/*
	  In: Body
	*/
	Payload *models.Server `json:"body,omitempty"`
}

// NewReplaceServerBackendOK creates ReplaceServerBackendOK with default headers values
func NewReplaceServerBackendOK() *ReplaceServerBackendOK {

	return &ReplaceServerBackendOK{}
}

// WithPayload adds the payload to the replace server backend o k response
func (o *ReplaceServerBackendOK) WithPayload(payload *models.Server) *ReplaceServerBackendOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server backend o k response
func (o *ReplaceServerBackendOK) SetPayload(payload *models.Server) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerBackendOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceServerBackendAcceptedCode is the HTTP code returned for type ReplaceServerBackendAccepted
const ReplaceServerBackendAcceptedCode int = 202

/*
ReplaceServerBackendAccepted Configuration change accepted and reload requested

swagger:response replaceServerBackendAccepted
*/
type ReplaceServerBackendAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Server `json:"body,omitempty"`
}

// NewReplaceServerBackendAccepted creates ReplaceServerBackendAccepted with default headers values
func NewReplaceServerBackendAccepted() *ReplaceServerBackendAccepted {

	return &ReplaceServerBackendAccepted{}
}

// WithReloadID adds the reloadId to the replace server backend accepted response
func (o *ReplaceServerBackendAccepted) WithReloadID(reloadID string) *ReplaceServerBackendAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace server backend accepted response
func (o *ReplaceServerBackendAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace server backend accepted response
func (o *ReplaceServerBackendAccepted) WithPayload(payload *models.Server) *ReplaceServerBackendAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server backend accepted response
func (o *ReplaceServerBackendAccepted) SetPayload(payload *models.Server) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerBackendAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceServerBackendBadRequestCode is the HTTP code returned for type ReplaceServerBackendBadRequest
const ReplaceServerBackendBadRequestCode int = 400

/*
ReplaceServerBackendBadRequest Bad request

swagger:response replaceServerBackendBadRequest
*/
type ReplaceServerBackendBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceServerBackendBadRequest creates ReplaceServerBackendBadRequest with default headers values
func NewReplaceServerBackendBadRequest() *ReplaceServerBackendBadRequest {

	return &ReplaceServerBackendBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace server backend bad request response
func (o *ReplaceServerBackendBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceServerBackendBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace server backend bad request response
func (o *ReplaceServerBackendBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace server backend bad request response
func (o *ReplaceServerBackendBadRequest) WithPayload(payload *models.Error) *ReplaceServerBackendBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server backend bad request response
func (o *ReplaceServerBackendBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerBackendBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceServerBackendNotFoundCode is the HTTP code returned for type ReplaceServerBackendNotFound
const ReplaceServerBackendNotFoundCode int = 404

/*
ReplaceServerBackendNotFound The specified resource was not found

swagger:response replaceServerBackendNotFound
*/
type ReplaceServerBackendNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceServerBackendNotFound creates ReplaceServerBackendNotFound with default headers values
func NewReplaceServerBackendNotFound() *ReplaceServerBackendNotFound {

	return &ReplaceServerBackendNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the replace server backend not found response
func (o *ReplaceServerBackendNotFound) WithConfigurationVersion(configurationVersion string) *ReplaceServerBackendNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace server backend not found response
func (o *ReplaceServerBackendNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace server backend not found response
func (o *ReplaceServerBackendNotFound) WithPayload(payload *models.Error) *ReplaceServerBackendNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server backend not found response
func (o *ReplaceServerBackendNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerBackendNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
ReplaceServerBackendDefault General Error

swagger:response replaceServerBackendDefault
*/
type ReplaceServerBackendDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceServerBackendDefault creates ReplaceServerBackendDefault with default headers values
func NewReplaceServerBackendDefault(code int) *ReplaceServerBackendDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceServerBackendDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace server backend default response
func (o *ReplaceServerBackendDefault) WithStatusCode(code int) *ReplaceServerBackendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace server backend default response
func (o *ReplaceServerBackendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace server backend default response
func (o *ReplaceServerBackendDefault) WithConfigurationVersion(configurationVersion string) *ReplaceServerBackendDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace server backend default response
func (o *ReplaceServerBackendDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace server backend default response
func (o *ReplaceServerBackendDefault) WithPayload(payload *models.Error) *ReplaceServerBackendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server backend default response
func (o *ReplaceServerBackendDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerBackendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
