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

package server_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// ReplaceServerSwitchingRulesOKCode is the HTTP code returned for type ReplaceServerSwitchingRulesOK
const ReplaceServerSwitchingRulesOKCode int = 200

/*
ReplaceServerSwitchingRulesOK All Server Switching Rule lines replaced

swagger:response replaceServerSwitchingRulesOK
*/
type ReplaceServerSwitchingRulesOK struct {

	/*
	  In: Body
	*/
	Payload models.ServerSwitchingRules `json:"body,omitempty"`
}

// NewReplaceServerSwitchingRulesOK creates ReplaceServerSwitchingRulesOK with default headers values
func NewReplaceServerSwitchingRulesOK() *ReplaceServerSwitchingRulesOK {

	return &ReplaceServerSwitchingRulesOK{}
}

// WithPayload adds the payload to the replace server switching rules o k response
func (o *ReplaceServerSwitchingRulesOK) WithPayload(payload models.ServerSwitchingRules) *ReplaceServerSwitchingRulesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server switching rules o k response
func (o *ReplaceServerSwitchingRulesOK) SetPayload(payload models.ServerSwitchingRules) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerSwitchingRulesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.ServerSwitchingRules{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ReplaceServerSwitchingRulesAcceptedCode is the HTTP code returned for type ReplaceServerSwitchingRulesAccepted
const ReplaceServerSwitchingRulesAcceptedCode int = 202

/*
ReplaceServerSwitchingRulesAccepted Configuration change accepted and reload requested

swagger:response replaceServerSwitchingRulesAccepted
*/
type ReplaceServerSwitchingRulesAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload models.ServerSwitchingRules `json:"body,omitempty"`
}

// NewReplaceServerSwitchingRulesAccepted creates ReplaceServerSwitchingRulesAccepted with default headers values
func NewReplaceServerSwitchingRulesAccepted() *ReplaceServerSwitchingRulesAccepted {

	return &ReplaceServerSwitchingRulesAccepted{}
}

// WithReloadID adds the reloadId to the replace server switching rules accepted response
func (o *ReplaceServerSwitchingRulesAccepted) WithReloadID(reloadID string) *ReplaceServerSwitchingRulesAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace server switching rules accepted response
func (o *ReplaceServerSwitchingRulesAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace server switching rules accepted response
func (o *ReplaceServerSwitchingRulesAccepted) WithPayload(payload models.ServerSwitchingRules) *ReplaceServerSwitchingRulesAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server switching rules accepted response
func (o *ReplaceServerSwitchingRulesAccepted) SetPayload(payload models.ServerSwitchingRules) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerSwitchingRulesAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.ServerSwitchingRules{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ReplaceServerSwitchingRulesBadRequestCode is the HTTP code returned for type ReplaceServerSwitchingRulesBadRequest
const ReplaceServerSwitchingRulesBadRequestCode int = 400

/*
ReplaceServerSwitchingRulesBadRequest Bad request

swagger:response replaceServerSwitchingRulesBadRequest
*/
type ReplaceServerSwitchingRulesBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceServerSwitchingRulesBadRequest creates ReplaceServerSwitchingRulesBadRequest with default headers values
func NewReplaceServerSwitchingRulesBadRequest() *ReplaceServerSwitchingRulesBadRequest {

	return &ReplaceServerSwitchingRulesBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace server switching rules bad request response
func (o *ReplaceServerSwitchingRulesBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceServerSwitchingRulesBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace server switching rules bad request response
func (o *ReplaceServerSwitchingRulesBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace server switching rules bad request response
func (o *ReplaceServerSwitchingRulesBadRequest) WithPayload(payload *models.Error) *ReplaceServerSwitchingRulesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server switching rules bad request response
func (o *ReplaceServerSwitchingRulesBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerSwitchingRulesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
ReplaceServerSwitchingRulesDefault General Error

swagger:response replaceServerSwitchingRulesDefault
*/
type ReplaceServerSwitchingRulesDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceServerSwitchingRulesDefault creates ReplaceServerSwitchingRulesDefault with default headers values
func NewReplaceServerSwitchingRulesDefault(code int) *ReplaceServerSwitchingRulesDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceServerSwitchingRulesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace server switching rules default response
func (o *ReplaceServerSwitchingRulesDefault) WithStatusCode(code int) *ReplaceServerSwitchingRulesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace server switching rules default response
func (o *ReplaceServerSwitchingRulesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace server switching rules default response
func (o *ReplaceServerSwitchingRulesDefault) WithConfigurationVersion(configurationVersion string) *ReplaceServerSwitchingRulesDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace server switching rules default response
func (o *ReplaceServerSwitchingRulesDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace server switching rules default response
func (o *ReplaceServerSwitchingRulesDefault) WithPayload(payload *models.Error) *ReplaceServerSwitchingRulesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server switching rules default response
func (o *ReplaceServerSwitchingRulesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerSwitchingRulesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
