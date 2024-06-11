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

package log_target

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// ReplaceLogTargetsOKCode is the HTTP code returned for type ReplaceLogTargetsOK
const ReplaceLogTargetsOKCode int = 200

/*
ReplaceLogTargetsOK All Log Target lines replaced

swagger:response replaceLogTargetsOK
*/
type ReplaceLogTargetsOK struct {

	/*
	  In: Body
	*/
	Payload models.LogTargets `json:"body,omitempty"`
}

// NewReplaceLogTargetsOK creates ReplaceLogTargetsOK with default headers values
func NewReplaceLogTargetsOK() *ReplaceLogTargetsOK {

	return &ReplaceLogTargetsOK{}
}

// WithPayload adds the payload to the replace log targets o k response
func (o *ReplaceLogTargetsOK) WithPayload(payload models.LogTargets) *ReplaceLogTargetsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace log targets o k response
func (o *ReplaceLogTargetsOK) SetPayload(payload models.LogTargets) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceLogTargetsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.LogTargets{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ReplaceLogTargetsAcceptedCode is the HTTP code returned for type ReplaceLogTargetsAccepted
const ReplaceLogTargetsAcceptedCode int = 202

/*
ReplaceLogTargetsAccepted Configuration change accepted and reload requested

swagger:response replaceLogTargetsAccepted
*/
type ReplaceLogTargetsAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload models.LogTargets `json:"body,omitempty"`
}

// NewReplaceLogTargetsAccepted creates ReplaceLogTargetsAccepted with default headers values
func NewReplaceLogTargetsAccepted() *ReplaceLogTargetsAccepted {

	return &ReplaceLogTargetsAccepted{}
}

// WithReloadID adds the reloadId to the replace log targets accepted response
func (o *ReplaceLogTargetsAccepted) WithReloadID(reloadID string) *ReplaceLogTargetsAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace log targets accepted response
func (o *ReplaceLogTargetsAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace log targets accepted response
func (o *ReplaceLogTargetsAccepted) WithPayload(payload models.LogTargets) *ReplaceLogTargetsAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace log targets accepted response
func (o *ReplaceLogTargetsAccepted) SetPayload(payload models.LogTargets) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceLogTargetsAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.LogTargets{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ReplaceLogTargetsBadRequestCode is the HTTP code returned for type ReplaceLogTargetsBadRequest
const ReplaceLogTargetsBadRequestCode int = 400

/*
ReplaceLogTargetsBadRequest Bad request

swagger:response replaceLogTargetsBadRequest
*/
type ReplaceLogTargetsBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceLogTargetsBadRequest creates ReplaceLogTargetsBadRequest with default headers values
func NewReplaceLogTargetsBadRequest() *ReplaceLogTargetsBadRequest {

	return &ReplaceLogTargetsBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace log targets bad request response
func (o *ReplaceLogTargetsBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceLogTargetsBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace log targets bad request response
func (o *ReplaceLogTargetsBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace log targets bad request response
func (o *ReplaceLogTargetsBadRequest) WithPayload(payload *models.Error) *ReplaceLogTargetsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace log targets bad request response
func (o *ReplaceLogTargetsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceLogTargetsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
ReplaceLogTargetsDefault General Error

swagger:response replaceLogTargetsDefault
*/
type ReplaceLogTargetsDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceLogTargetsDefault creates ReplaceLogTargetsDefault with default headers values
func NewReplaceLogTargetsDefault(code int) *ReplaceLogTargetsDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceLogTargetsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace log targets default response
func (o *ReplaceLogTargetsDefault) WithStatusCode(code int) *ReplaceLogTargetsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace log targets default response
func (o *ReplaceLogTargetsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace log targets default response
func (o *ReplaceLogTargetsDefault) WithConfigurationVersion(configurationVersion string) *ReplaceLogTargetsDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace log targets default response
func (o *ReplaceLogTargetsDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace log targets default response
func (o *ReplaceLogTargetsDefault) WithPayload(payload *models.Error) *ReplaceLogTargetsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace log targets default response
func (o *ReplaceLogTargetsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceLogTargetsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
