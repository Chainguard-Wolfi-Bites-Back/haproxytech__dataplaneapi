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

package frontend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models/v2"
)

// ReplaceFrontendOKCode is the HTTP code returned for type ReplaceFrontendOK
const ReplaceFrontendOKCode int = 200

/*ReplaceFrontendOK Frontend replaced

swagger:response replaceFrontendOK
*/
type ReplaceFrontendOK struct {

	/*
	  In: Body
	*/
	Payload *models.Frontend `json:"body,omitempty"`
}

// NewReplaceFrontendOK creates ReplaceFrontendOK with default headers values
func NewReplaceFrontendOK() *ReplaceFrontendOK {

	return &ReplaceFrontendOK{}
}

// WithPayload adds the payload to the replace frontend o k response
func (o *ReplaceFrontendOK) WithPayload(payload *models.Frontend) *ReplaceFrontendOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace frontend o k response
func (o *ReplaceFrontendOK) SetPayload(payload *models.Frontend) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceFrontendOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceFrontendAcceptedCode is the HTTP code returned for type ReplaceFrontendAccepted
const ReplaceFrontendAcceptedCode int = 202

/*ReplaceFrontendAccepted Configuration change accepted and reload requested

swagger:response replaceFrontendAccepted
*/
type ReplaceFrontendAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Frontend `json:"body,omitempty"`
}

// NewReplaceFrontendAccepted creates ReplaceFrontendAccepted with default headers values
func NewReplaceFrontendAccepted() *ReplaceFrontendAccepted {

	return &ReplaceFrontendAccepted{}
}

// WithReloadID adds the reloadId to the replace frontend accepted response
func (o *ReplaceFrontendAccepted) WithReloadID(reloadID string) *ReplaceFrontendAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace frontend accepted response
func (o *ReplaceFrontendAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace frontend accepted response
func (o *ReplaceFrontendAccepted) WithPayload(payload *models.Frontend) *ReplaceFrontendAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace frontend accepted response
func (o *ReplaceFrontendAccepted) SetPayload(payload *models.Frontend) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceFrontendAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceFrontendBadRequestCode is the HTTP code returned for type ReplaceFrontendBadRequest
const ReplaceFrontendBadRequestCode int = 400

/*ReplaceFrontendBadRequest Bad request

swagger:response replaceFrontendBadRequest
*/
type ReplaceFrontendBadRequest struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceFrontendBadRequest creates ReplaceFrontendBadRequest with default headers values
func NewReplaceFrontendBadRequest() *ReplaceFrontendBadRequest {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &ReplaceFrontendBadRequest{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the replace frontend bad request response
func (o *ReplaceFrontendBadRequest) WithConfigurationVersion(configurationVersion int64) *ReplaceFrontendBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace frontend bad request response
func (o *ReplaceFrontendBadRequest) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace frontend bad request response
func (o *ReplaceFrontendBadRequest) WithPayload(payload *models.Error) *ReplaceFrontendBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace frontend bad request response
func (o *ReplaceFrontendBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceFrontendBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceFrontendNotFoundCode is the HTTP code returned for type ReplaceFrontendNotFound
const ReplaceFrontendNotFoundCode int = 404

/*ReplaceFrontendNotFound The specified resource was not found

swagger:response replaceFrontendNotFound
*/
type ReplaceFrontendNotFound struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceFrontendNotFound creates ReplaceFrontendNotFound with default headers values
func NewReplaceFrontendNotFound() *ReplaceFrontendNotFound {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &ReplaceFrontendNotFound{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the replace frontend not found response
func (o *ReplaceFrontendNotFound) WithConfigurationVersion(configurationVersion int64) *ReplaceFrontendNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace frontend not found response
func (o *ReplaceFrontendNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace frontend not found response
func (o *ReplaceFrontendNotFound) WithPayload(payload *models.Error) *ReplaceFrontendNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace frontend not found response
func (o *ReplaceFrontendNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceFrontendNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
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

/*ReplaceFrontendDefault General Error

swagger:response replaceFrontendDefault
*/
type ReplaceFrontendDefault struct {
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

// NewReplaceFrontendDefault creates ReplaceFrontendDefault with default headers values
func NewReplaceFrontendDefault(code int) *ReplaceFrontendDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &ReplaceFrontendDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the replace frontend default response
func (o *ReplaceFrontendDefault) WithStatusCode(code int) *ReplaceFrontendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace frontend default response
func (o *ReplaceFrontendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace frontend default response
func (o *ReplaceFrontendDefault) WithConfigurationVersion(configurationVersion int64) *ReplaceFrontendDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace frontend default response
func (o *ReplaceFrontendDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace frontend default response
func (o *ReplaceFrontendDefault) WithPayload(payload *models.Error) *ReplaceFrontendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace frontend default response
func (o *ReplaceFrontendDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceFrontendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
