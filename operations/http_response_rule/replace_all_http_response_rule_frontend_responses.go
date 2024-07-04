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

package http_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// ReplaceAllHTTPResponseRuleFrontendOKCode is the HTTP code returned for type ReplaceAllHTTPResponseRuleFrontendOK
const ReplaceAllHTTPResponseRuleFrontendOKCode int = 200

/*
ReplaceAllHTTPResponseRuleFrontendOK All HTTP Response Rule lines replaced

swagger:response replaceAllHttpResponseRuleFrontendOK
*/
type ReplaceAllHTTPResponseRuleFrontendOK struct {

	/*
	  In: Body
	*/
	Payload models.HTTPResponseRules `json:"body,omitempty"`
}

// NewReplaceAllHTTPResponseRuleFrontendOK creates ReplaceAllHTTPResponseRuleFrontendOK with default headers values
func NewReplaceAllHTTPResponseRuleFrontendOK() *ReplaceAllHTTPResponseRuleFrontendOK {

	return &ReplaceAllHTTPResponseRuleFrontendOK{}
}

// WithPayload adds the payload to the replace all Http response rule frontend o k response
func (o *ReplaceAllHTTPResponseRuleFrontendOK) WithPayload(payload models.HTTPResponseRules) *ReplaceAllHTTPResponseRuleFrontendOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace all Http response rule frontend o k response
func (o *ReplaceAllHTTPResponseRuleFrontendOK) SetPayload(payload models.HTTPResponseRules) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAllHTTPResponseRuleFrontendOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.HTTPResponseRules{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ReplaceAllHTTPResponseRuleFrontendAcceptedCode is the HTTP code returned for type ReplaceAllHTTPResponseRuleFrontendAccepted
const ReplaceAllHTTPResponseRuleFrontendAcceptedCode int = 202

/*
ReplaceAllHTTPResponseRuleFrontendAccepted Configuration change accepted and reload requested

swagger:response replaceAllHttpResponseRuleFrontendAccepted
*/
type ReplaceAllHTTPResponseRuleFrontendAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload models.HTTPResponseRules `json:"body,omitempty"`
}

// NewReplaceAllHTTPResponseRuleFrontendAccepted creates ReplaceAllHTTPResponseRuleFrontendAccepted with default headers values
func NewReplaceAllHTTPResponseRuleFrontendAccepted() *ReplaceAllHTTPResponseRuleFrontendAccepted {

	return &ReplaceAllHTTPResponseRuleFrontendAccepted{}
}

// WithReloadID adds the reloadId to the replace all Http response rule frontend accepted response
func (o *ReplaceAllHTTPResponseRuleFrontendAccepted) WithReloadID(reloadID string) *ReplaceAllHTTPResponseRuleFrontendAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace all Http response rule frontend accepted response
func (o *ReplaceAllHTTPResponseRuleFrontendAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace all Http response rule frontend accepted response
func (o *ReplaceAllHTTPResponseRuleFrontendAccepted) WithPayload(payload models.HTTPResponseRules) *ReplaceAllHTTPResponseRuleFrontendAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace all Http response rule frontend accepted response
func (o *ReplaceAllHTTPResponseRuleFrontendAccepted) SetPayload(payload models.HTTPResponseRules) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAllHTTPResponseRuleFrontendAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.HTTPResponseRules{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ReplaceAllHTTPResponseRuleFrontendBadRequestCode is the HTTP code returned for type ReplaceAllHTTPResponseRuleFrontendBadRequest
const ReplaceAllHTTPResponseRuleFrontendBadRequestCode int = 400

/*
ReplaceAllHTTPResponseRuleFrontendBadRequest Bad request

swagger:response replaceAllHttpResponseRuleFrontendBadRequest
*/
type ReplaceAllHTTPResponseRuleFrontendBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceAllHTTPResponseRuleFrontendBadRequest creates ReplaceAllHTTPResponseRuleFrontendBadRequest with default headers values
func NewReplaceAllHTTPResponseRuleFrontendBadRequest() *ReplaceAllHTTPResponseRuleFrontendBadRequest {

	return &ReplaceAllHTTPResponseRuleFrontendBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace all Http response rule frontend bad request response
func (o *ReplaceAllHTTPResponseRuleFrontendBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceAllHTTPResponseRuleFrontendBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace all Http response rule frontend bad request response
func (o *ReplaceAllHTTPResponseRuleFrontendBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace all Http response rule frontend bad request response
func (o *ReplaceAllHTTPResponseRuleFrontendBadRequest) WithPayload(payload *models.Error) *ReplaceAllHTTPResponseRuleFrontendBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace all Http response rule frontend bad request response
func (o *ReplaceAllHTTPResponseRuleFrontendBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAllHTTPResponseRuleFrontendBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
ReplaceAllHTTPResponseRuleFrontendDefault General Error

swagger:response replaceAllHttpResponseRuleFrontendDefault
*/
type ReplaceAllHTTPResponseRuleFrontendDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceAllHTTPResponseRuleFrontendDefault creates ReplaceAllHTTPResponseRuleFrontendDefault with default headers values
func NewReplaceAllHTTPResponseRuleFrontendDefault(code int) *ReplaceAllHTTPResponseRuleFrontendDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceAllHTTPResponseRuleFrontendDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace all HTTP response rule frontend default response
func (o *ReplaceAllHTTPResponseRuleFrontendDefault) WithStatusCode(code int) *ReplaceAllHTTPResponseRuleFrontendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace all HTTP response rule frontend default response
func (o *ReplaceAllHTTPResponseRuleFrontendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace all HTTP response rule frontend default response
func (o *ReplaceAllHTTPResponseRuleFrontendDefault) WithConfigurationVersion(configurationVersion string) *ReplaceAllHTTPResponseRuleFrontendDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace all HTTP response rule frontend default response
func (o *ReplaceAllHTTPResponseRuleFrontendDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace all HTTP response rule frontend default response
func (o *ReplaceAllHTTPResponseRuleFrontendDefault) WithPayload(payload *models.Error) *ReplaceAllHTTPResponseRuleFrontendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace all HTTP response rule frontend default response
func (o *ReplaceAllHTTPResponseRuleFrontendDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAllHTTPResponseRuleFrontendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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