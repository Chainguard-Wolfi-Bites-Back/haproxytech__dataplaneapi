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

package tcp_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models/v2"
)

// CreateTCPResponseRuleCreatedCode is the HTTP code returned for type CreateTCPResponseRuleCreated
const CreateTCPResponseRuleCreatedCode int = 201

/*CreateTCPResponseRuleCreated TCP Response Rule created

swagger:response createTcpResponseRuleCreated
*/
type CreateTCPResponseRuleCreated struct {

	/*
	  In: Body
	*/
	Payload *models.TCPResponseRule `json:"body,omitempty"`
}

// NewCreateTCPResponseRuleCreated creates CreateTCPResponseRuleCreated with default headers values
func NewCreateTCPResponseRuleCreated() *CreateTCPResponseRuleCreated {

	return &CreateTCPResponseRuleCreated{}
}

// WithPayload adds the payload to the create Tcp response rule created response
func (o *CreateTCPResponseRuleCreated) WithPayload(payload *models.TCPResponseRule) *CreateTCPResponseRuleCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Tcp response rule created response
func (o *CreateTCPResponseRuleCreated) SetPayload(payload *models.TCPResponseRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTCPResponseRuleCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateTCPResponseRuleAcceptedCode is the HTTP code returned for type CreateTCPResponseRuleAccepted
const CreateTCPResponseRuleAcceptedCode int = 202

/*CreateTCPResponseRuleAccepted Configuration change accepted and reload requested

swagger:response createTcpResponseRuleAccepted
*/
type CreateTCPResponseRuleAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.TCPResponseRule `json:"body,omitempty"`
}

// NewCreateTCPResponseRuleAccepted creates CreateTCPResponseRuleAccepted with default headers values
func NewCreateTCPResponseRuleAccepted() *CreateTCPResponseRuleAccepted {

	return &CreateTCPResponseRuleAccepted{}
}

// WithReloadID adds the reloadId to the create Tcp response rule accepted response
func (o *CreateTCPResponseRuleAccepted) WithReloadID(reloadID string) *CreateTCPResponseRuleAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create Tcp response rule accepted response
func (o *CreateTCPResponseRuleAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create Tcp response rule accepted response
func (o *CreateTCPResponseRuleAccepted) WithPayload(payload *models.TCPResponseRule) *CreateTCPResponseRuleAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Tcp response rule accepted response
func (o *CreateTCPResponseRuleAccepted) SetPayload(payload *models.TCPResponseRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTCPResponseRuleAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateTCPResponseRuleBadRequestCode is the HTTP code returned for type CreateTCPResponseRuleBadRequest
const CreateTCPResponseRuleBadRequestCode int = 400

/*CreateTCPResponseRuleBadRequest Bad request

swagger:response createTcpResponseRuleBadRequest
*/
type CreateTCPResponseRuleBadRequest struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateTCPResponseRuleBadRequest creates CreateTCPResponseRuleBadRequest with default headers values
func NewCreateTCPResponseRuleBadRequest() *CreateTCPResponseRuleBadRequest {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CreateTCPResponseRuleBadRequest{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the create Tcp response rule bad request response
func (o *CreateTCPResponseRuleBadRequest) WithConfigurationVersion(configurationVersion int64) *CreateTCPResponseRuleBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create Tcp response rule bad request response
func (o *CreateTCPResponseRuleBadRequest) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create Tcp response rule bad request response
func (o *CreateTCPResponseRuleBadRequest) WithPayload(payload *models.Error) *CreateTCPResponseRuleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Tcp response rule bad request response
func (o *CreateTCPResponseRuleBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTCPResponseRuleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
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

// CreateTCPResponseRuleConflictCode is the HTTP code returned for type CreateTCPResponseRuleConflict
const CreateTCPResponseRuleConflictCode int = 409

/*CreateTCPResponseRuleConflict The specified resource already exists

swagger:response createTcpResponseRuleConflict
*/
type CreateTCPResponseRuleConflict struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateTCPResponseRuleConflict creates CreateTCPResponseRuleConflict with default headers values
func NewCreateTCPResponseRuleConflict() *CreateTCPResponseRuleConflict {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CreateTCPResponseRuleConflict{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the create Tcp response rule conflict response
func (o *CreateTCPResponseRuleConflict) WithConfigurationVersion(configurationVersion int64) *CreateTCPResponseRuleConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create Tcp response rule conflict response
func (o *CreateTCPResponseRuleConflict) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create Tcp response rule conflict response
func (o *CreateTCPResponseRuleConflict) WithPayload(payload *models.Error) *CreateTCPResponseRuleConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Tcp response rule conflict response
func (o *CreateTCPResponseRuleConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTCPResponseRuleConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
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

/*CreateTCPResponseRuleDefault General Error

swagger:response createTcpResponseRuleDefault
*/
type CreateTCPResponseRuleDefault struct {
	_statusCode int
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateTCPResponseRuleDefault creates CreateTCPResponseRuleDefault with default headers values
func NewCreateTCPResponseRuleDefault(code int) *CreateTCPResponseRuleDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CreateTCPResponseRuleDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the create TCP response rule default response
func (o *CreateTCPResponseRuleDefault) WithStatusCode(code int) *CreateTCPResponseRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create TCP response rule default response
func (o *CreateTCPResponseRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create TCP response rule default response
func (o *CreateTCPResponseRuleDefault) WithConfigurationVersion(configurationVersion int64) *CreateTCPResponseRuleDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create TCP response rule default response
func (o *CreateTCPResponseRuleDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create TCP response rule default response
func (o *CreateTCPResponseRuleDefault) WithPayload(payload *models.Error) *CreateTCPResponseRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create TCP response rule default response
func (o *CreateTCPResponseRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTCPResponseRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
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
