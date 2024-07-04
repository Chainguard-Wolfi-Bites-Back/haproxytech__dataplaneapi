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

package http_error_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// GetHTTPErrorRuleFrontendOKCode is the HTTP code returned for type GetHTTPErrorRuleFrontendOK
const GetHTTPErrorRuleFrontendOKCode int = 200

/*
GetHTTPErrorRuleFrontendOK Successful operation

swagger:response getHttpErrorRuleFrontendOK
*/
type GetHTTPErrorRuleFrontendOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.HTTPErrorRule `json:"body,omitempty"`
}

// NewGetHTTPErrorRuleFrontendOK creates GetHTTPErrorRuleFrontendOK with default headers values
func NewGetHTTPErrorRuleFrontendOK() *GetHTTPErrorRuleFrontendOK {

	return &GetHTTPErrorRuleFrontendOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get Http error rule frontend o k response
func (o *GetHTTPErrorRuleFrontendOK) WithConfigurationVersion(configurationVersion string) *GetHTTPErrorRuleFrontendOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Http error rule frontend o k response
func (o *GetHTTPErrorRuleFrontendOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Http error rule frontend o k response
func (o *GetHTTPErrorRuleFrontendOK) WithPayload(payload *models.HTTPErrorRule) *GetHTTPErrorRuleFrontendOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Http error rule frontend o k response
func (o *GetHTTPErrorRuleFrontendOK) SetPayload(payload *models.HTTPErrorRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHTTPErrorRuleFrontendOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHTTPErrorRuleFrontendNotFoundCode is the HTTP code returned for type GetHTTPErrorRuleFrontendNotFound
const GetHTTPErrorRuleFrontendNotFoundCode int = 404

/*
GetHTTPErrorRuleFrontendNotFound The specified resource was not found

swagger:response getHttpErrorRuleFrontendNotFound
*/
type GetHTTPErrorRuleFrontendNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHTTPErrorRuleFrontendNotFound creates GetHTTPErrorRuleFrontendNotFound with default headers values
func NewGetHTTPErrorRuleFrontendNotFound() *GetHTTPErrorRuleFrontendNotFound {

	return &GetHTTPErrorRuleFrontendNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get Http error rule frontend not found response
func (o *GetHTTPErrorRuleFrontendNotFound) WithConfigurationVersion(configurationVersion string) *GetHTTPErrorRuleFrontendNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Http error rule frontend not found response
func (o *GetHTTPErrorRuleFrontendNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Http error rule frontend not found response
func (o *GetHTTPErrorRuleFrontendNotFound) WithPayload(payload *models.Error) *GetHTTPErrorRuleFrontendNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Http error rule frontend not found response
func (o *GetHTTPErrorRuleFrontendNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHTTPErrorRuleFrontendNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
GetHTTPErrorRuleFrontendDefault General Error

swagger:response getHttpErrorRuleFrontendDefault
*/
type GetHTTPErrorRuleFrontendDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHTTPErrorRuleFrontendDefault creates GetHTTPErrorRuleFrontendDefault with default headers values
func NewGetHTTPErrorRuleFrontendDefault(code int) *GetHTTPErrorRuleFrontendDefault {
	if code <= 0 {
		code = 500
	}

	return &GetHTTPErrorRuleFrontendDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get HTTP error rule frontend default response
func (o *GetHTTPErrorRuleFrontendDefault) WithStatusCode(code int) *GetHTTPErrorRuleFrontendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get HTTP error rule frontend default response
func (o *GetHTTPErrorRuleFrontendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get HTTP error rule frontend default response
func (o *GetHTTPErrorRuleFrontendDefault) WithConfigurationVersion(configurationVersion string) *GetHTTPErrorRuleFrontendDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get HTTP error rule frontend default response
func (o *GetHTTPErrorRuleFrontendDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get HTTP error rule frontend default response
func (o *GetHTTPErrorRuleFrontendDefault) WithPayload(payload *models.Error) *GetHTTPErrorRuleFrontendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get HTTP error rule frontend default response
func (o *GetHTTPErrorRuleFrontendDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHTTPErrorRuleFrontendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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