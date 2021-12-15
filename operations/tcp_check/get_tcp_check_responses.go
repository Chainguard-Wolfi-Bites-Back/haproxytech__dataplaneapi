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

package tcp_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v2/models"
)

// GetTCPCheckOKCode is the HTTP code returned for type GetTCPCheckOK
const GetTCPCheckOKCode int = 200

/*GetTCPCheckOK Successful operation

swagger:response getTcpCheckOK
*/
type GetTCPCheckOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetTCPCheckOKBody `json:"body,omitempty"`
}

// NewGetTCPCheckOK creates GetTCPCheckOK with default headers values
func NewGetTCPCheckOK() *GetTCPCheckOK {

	return &GetTCPCheckOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get Tcp check o k response
func (o *GetTCPCheckOK) WithConfigurationVersion(configurationVersion string) *GetTCPCheckOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Tcp check o k response
func (o *GetTCPCheckOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Tcp check o k response
func (o *GetTCPCheckOK) WithPayload(payload *GetTCPCheckOKBody) *GetTCPCheckOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Tcp check o k response
func (o *GetTCPCheckOK) SetPayload(payload *GetTCPCheckOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTCPCheckOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetTCPCheckNotFoundCode is the HTTP code returned for type GetTCPCheckNotFound
const GetTCPCheckNotFoundCode int = 404

/*GetTCPCheckNotFound The specified resource was not found

swagger:response getTcpCheckNotFound
*/
type GetTCPCheckNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTCPCheckNotFound creates GetTCPCheckNotFound with default headers values
func NewGetTCPCheckNotFound() *GetTCPCheckNotFound {

	return &GetTCPCheckNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get Tcp check not found response
func (o *GetTCPCheckNotFound) WithConfigurationVersion(configurationVersion string) *GetTCPCheckNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Tcp check not found response
func (o *GetTCPCheckNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Tcp check not found response
func (o *GetTCPCheckNotFound) WithPayload(payload *models.Error) *GetTCPCheckNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Tcp check not found response
func (o *GetTCPCheckNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTCPCheckNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetTCPCheckDefault General Error

swagger:response getTcpCheckDefault
*/
type GetTCPCheckDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTCPCheckDefault creates GetTCPCheckDefault with default headers values
func NewGetTCPCheckDefault(code int) *GetTCPCheckDefault {
	if code <= 0 {
		code = 500
	}

	return &GetTCPCheckDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get TCP check default response
func (o *GetTCPCheckDefault) WithStatusCode(code int) *GetTCPCheckDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get TCP check default response
func (o *GetTCPCheckDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get TCP check default response
func (o *GetTCPCheckDefault) WithConfigurationVersion(configurationVersion string) *GetTCPCheckDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get TCP check default response
func (o *GetTCPCheckDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get TCP check default response
func (o *GetTCPCheckDefault) WithPayload(payload *models.Error) *GetTCPCheckDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get TCP check default response
func (o *GetTCPCheckDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTCPCheckDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
