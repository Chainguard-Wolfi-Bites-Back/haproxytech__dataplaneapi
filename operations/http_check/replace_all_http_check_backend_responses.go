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

package http_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// ReplaceAllHTTPCheckBackendOKCode is the HTTP code returned for type ReplaceAllHTTPCheckBackendOK
const ReplaceAllHTTPCheckBackendOKCode int = 200

/*
ReplaceAllHTTPCheckBackendOK All HTTP checks lines replaced

swagger:response replaceAllHttpCheckBackendOK
*/
type ReplaceAllHTTPCheckBackendOK struct {

	/*
	  In: Body
	*/
	Payload models.HTTPChecks `json:"body,omitempty"`
}

// NewReplaceAllHTTPCheckBackendOK creates ReplaceAllHTTPCheckBackendOK with default headers values
func NewReplaceAllHTTPCheckBackendOK() *ReplaceAllHTTPCheckBackendOK {

	return &ReplaceAllHTTPCheckBackendOK{}
}

// WithPayload adds the payload to the replace all Http check backend o k response
func (o *ReplaceAllHTTPCheckBackendOK) WithPayload(payload models.HTTPChecks) *ReplaceAllHTTPCheckBackendOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace all Http check backend o k response
func (o *ReplaceAllHTTPCheckBackendOK) SetPayload(payload models.HTTPChecks) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAllHTTPCheckBackendOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.HTTPChecks{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ReplaceAllHTTPCheckBackendAcceptedCode is the HTTP code returned for type ReplaceAllHTTPCheckBackendAccepted
const ReplaceAllHTTPCheckBackendAcceptedCode int = 202

/*
ReplaceAllHTTPCheckBackendAccepted Configuration change accepted and reload requested

swagger:response replaceAllHttpCheckBackendAccepted
*/
type ReplaceAllHTTPCheckBackendAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload models.HTTPChecks `json:"body,omitempty"`
}

// NewReplaceAllHTTPCheckBackendAccepted creates ReplaceAllHTTPCheckBackendAccepted with default headers values
func NewReplaceAllHTTPCheckBackendAccepted() *ReplaceAllHTTPCheckBackendAccepted {

	return &ReplaceAllHTTPCheckBackendAccepted{}
}

// WithReloadID adds the reloadId to the replace all Http check backend accepted response
func (o *ReplaceAllHTTPCheckBackendAccepted) WithReloadID(reloadID string) *ReplaceAllHTTPCheckBackendAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace all Http check backend accepted response
func (o *ReplaceAllHTTPCheckBackendAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace all Http check backend accepted response
func (o *ReplaceAllHTTPCheckBackendAccepted) WithPayload(payload models.HTTPChecks) *ReplaceAllHTTPCheckBackendAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace all Http check backend accepted response
func (o *ReplaceAllHTTPCheckBackendAccepted) SetPayload(payload models.HTTPChecks) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAllHTTPCheckBackendAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.HTTPChecks{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ReplaceAllHTTPCheckBackendBadRequestCode is the HTTP code returned for type ReplaceAllHTTPCheckBackendBadRequest
const ReplaceAllHTTPCheckBackendBadRequestCode int = 400

/*
ReplaceAllHTTPCheckBackendBadRequest Bad request

swagger:response replaceAllHttpCheckBackendBadRequest
*/
type ReplaceAllHTTPCheckBackendBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceAllHTTPCheckBackendBadRequest creates ReplaceAllHTTPCheckBackendBadRequest with default headers values
func NewReplaceAllHTTPCheckBackendBadRequest() *ReplaceAllHTTPCheckBackendBadRequest {

	return &ReplaceAllHTTPCheckBackendBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace all Http check backend bad request response
func (o *ReplaceAllHTTPCheckBackendBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceAllHTTPCheckBackendBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace all Http check backend bad request response
func (o *ReplaceAllHTTPCheckBackendBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace all Http check backend bad request response
func (o *ReplaceAllHTTPCheckBackendBadRequest) WithPayload(payload *models.Error) *ReplaceAllHTTPCheckBackendBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace all Http check backend bad request response
func (o *ReplaceAllHTTPCheckBackendBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAllHTTPCheckBackendBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
ReplaceAllHTTPCheckBackendDefault General Error

swagger:response replaceAllHttpCheckBackendDefault
*/
type ReplaceAllHTTPCheckBackendDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceAllHTTPCheckBackendDefault creates ReplaceAllHTTPCheckBackendDefault with default headers values
func NewReplaceAllHTTPCheckBackendDefault(code int) *ReplaceAllHTTPCheckBackendDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceAllHTTPCheckBackendDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace all HTTP check backend default response
func (o *ReplaceAllHTTPCheckBackendDefault) WithStatusCode(code int) *ReplaceAllHTTPCheckBackendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace all HTTP check backend default response
func (o *ReplaceAllHTTPCheckBackendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace all HTTP check backend default response
func (o *ReplaceAllHTTPCheckBackendDefault) WithConfigurationVersion(configurationVersion string) *ReplaceAllHTTPCheckBackendDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace all HTTP check backend default response
func (o *ReplaceAllHTTPCheckBackendDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace all HTTP check backend default response
func (o *ReplaceAllHTTPCheckBackendDefault) WithPayload(payload *models.Error) *ReplaceAllHTTPCheckBackendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace all HTTP check backend default response
func (o *ReplaceAllHTTPCheckBackendDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAllHTTPCheckBackendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
