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

package maps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models/v2"
)

// GetRuntimeMapEntryOKCode is the HTTP code returned for type GetRuntimeMapEntryOK
const GetRuntimeMapEntryOKCode int = 200

/*GetRuntimeMapEntryOK Successful operation

swagger:response getRuntimeMapEntryOK
*/
type GetRuntimeMapEntryOK struct {

	/*
	  In: Body
	*/
	Payload *models.MapEntry `json:"body,omitempty"`
}

// NewGetRuntimeMapEntryOK creates GetRuntimeMapEntryOK with default headers values
func NewGetRuntimeMapEntryOK() *GetRuntimeMapEntryOK {

	return &GetRuntimeMapEntryOK{}
}

// WithPayload adds the payload to the get runtime map entry o k response
func (o *GetRuntimeMapEntryOK) WithPayload(payload *models.MapEntry) *GetRuntimeMapEntryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get runtime map entry o k response
func (o *GetRuntimeMapEntryOK) SetPayload(payload *models.MapEntry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRuntimeMapEntryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRuntimeMapEntryNotFoundCode is the HTTP code returned for type GetRuntimeMapEntryNotFound
const GetRuntimeMapEntryNotFoundCode int = 404

/*GetRuntimeMapEntryNotFound The specified resource was not found

swagger:response getRuntimeMapEntryNotFound
*/
type GetRuntimeMapEntryNotFound struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRuntimeMapEntryNotFound creates GetRuntimeMapEntryNotFound with default headers values
func NewGetRuntimeMapEntryNotFound() *GetRuntimeMapEntryNotFound {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetRuntimeMapEntryNotFound{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the get runtime map entry not found response
func (o *GetRuntimeMapEntryNotFound) WithConfigurationVersion(configurationVersion int64) *GetRuntimeMapEntryNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get runtime map entry not found response
func (o *GetRuntimeMapEntryNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get runtime map entry not found response
func (o *GetRuntimeMapEntryNotFound) WithPayload(payload *models.Error) *GetRuntimeMapEntryNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get runtime map entry not found response
func (o *GetRuntimeMapEntryNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRuntimeMapEntryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetRuntimeMapEntryDefault General Error

swagger:response getRuntimeMapEntryDefault
*/
type GetRuntimeMapEntryDefault struct {
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

// NewGetRuntimeMapEntryDefault creates GetRuntimeMapEntryDefault with default headers values
func NewGetRuntimeMapEntryDefault(code int) *GetRuntimeMapEntryDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetRuntimeMapEntryDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the get runtime map entry default response
func (o *GetRuntimeMapEntryDefault) WithStatusCode(code int) *GetRuntimeMapEntryDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get runtime map entry default response
func (o *GetRuntimeMapEntryDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get runtime map entry default response
func (o *GetRuntimeMapEntryDefault) WithConfigurationVersion(configurationVersion int64) *GetRuntimeMapEntryDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get runtime map entry default response
func (o *GetRuntimeMapEntryDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get runtime map entry default response
func (o *GetRuntimeMapEntryDefault) WithPayload(payload *models.Error) *GetRuntimeMapEntryDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get runtime map entry default response
func (o *GetRuntimeMapEntryDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRuntimeMapEntryDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
