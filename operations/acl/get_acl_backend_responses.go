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

package acl

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// GetACLBackendOKCode is the HTTP code returned for type GetACLBackendOK
const GetACLBackendOKCode int = 200

/*
GetACLBackendOK Successful operation

swagger:response getAclBackendOK
*/
type GetACLBackendOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.ACL `json:"body,omitempty"`
}

// NewGetACLBackendOK creates GetACLBackendOK with default headers values
func NewGetACLBackendOK() *GetACLBackendOK {

	return &GetACLBackendOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get Acl backend o k response
func (o *GetACLBackendOK) WithConfigurationVersion(configurationVersion string) *GetACLBackendOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Acl backend o k response
func (o *GetACLBackendOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Acl backend o k response
func (o *GetACLBackendOK) WithPayload(payload *models.ACL) *GetACLBackendOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Acl backend o k response
func (o *GetACLBackendOK) SetPayload(payload *models.ACL) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetACLBackendOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetACLBackendNotFoundCode is the HTTP code returned for type GetACLBackendNotFound
const GetACLBackendNotFoundCode int = 404

/*
GetACLBackendNotFound The specified resource was not found

swagger:response getAclBackendNotFound
*/
type GetACLBackendNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetACLBackendNotFound creates GetACLBackendNotFound with default headers values
func NewGetACLBackendNotFound() *GetACLBackendNotFound {

	return &GetACLBackendNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get Acl backend not found response
func (o *GetACLBackendNotFound) WithConfigurationVersion(configurationVersion string) *GetACLBackendNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Acl backend not found response
func (o *GetACLBackendNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Acl backend not found response
func (o *GetACLBackendNotFound) WithPayload(payload *models.Error) *GetACLBackendNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Acl backend not found response
func (o *GetACLBackendNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetACLBackendNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
GetACLBackendDefault General Error

swagger:response getAclBackendDefault
*/
type GetACLBackendDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetACLBackendDefault creates GetACLBackendDefault with default headers values
func NewGetACLBackendDefault(code int) *GetACLBackendDefault {
	if code <= 0 {
		code = 500
	}

	return &GetACLBackendDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get Acl backend default response
func (o *GetACLBackendDefault) WithStatusCode(code int) *GetACLBackendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get Acl backend default response
func (o *GetACLBackendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get Acl backend default response
func (o *GetACLBackendDefault) WithConfigurationVersion(configurationVersion string) *GetACLBackendDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Acl backend default response
func (o *GetACLBackendDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Acl backend default response
func (o *GetACLBackendDefault) WithPayload(payload *models.Error) *GetACLBackendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Acl backend default response
func (o *GetACLBackendDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetACLBackendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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