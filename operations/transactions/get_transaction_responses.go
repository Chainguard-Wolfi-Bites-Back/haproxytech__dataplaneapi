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

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models/v2"
)

// GetTransactionOKCode is the HTTP code returned for type GetTransactionOK
const GetTransactionOKCode int = 200

/*GetTransactionOK Successful operation

swagger:response getTransactionOK
*/
type GetTransactionOK struct {

	/*
	  In: Body
	*/
	Payload *models.Transaction `json:"body,omitempty"`
}

// NewGetTransactionOK creates GetTransactionOK with default headers values
func NewGetTransactionOK() *GetTransactionOK {

	return &GetTransactionOK{}
}

// WithPayload adds the payload to the get transaction o k response
func (o *GetTransactionOK) WithPayload(payload *models.Transaction) *GetTransactionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get transaction o k response
func (o *GetTransactionOK) SetPayload(payload *models.Transaction) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTransactionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTransactionNotFoundCode is the HTTP code returned for type GetTransactionNotFound
const GetTransactionNotFoundCode int = 404

/*GetTransactionNotFound The specified resource was not found

swagger:response getTransactionNotFound
*/
type GetTransactionNotFound struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTransactionNotFound creates GetTransactionNotFound with default headers values
func NewGetTransactionNotFound() *GetTransactionNotFound {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetTransactionNotFound{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the get transaction not found response
func (o *GetTransactionNotFound) WithConfigurationVersion(configurationVersion int64) *GetTransactionNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get transaction not found response
func (o *GetTransactionNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get transaction not found response
func (o *GetTransactionNotFound) WithPayload(payload *models.Error) *GetTransactionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get transaction not found response
func (o *GetTransactionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTransactionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetTransactionDefault General Error

swagger:response getTransactionDefault
*/
type GetTransactionDefault struct {
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

// NewGetTransactionDefault creates GetTransactionDefault with default headers values
func NewGetTransactionDefault(code int) *GetTransactionDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetTransactionDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the get transaction default response
func (o *GetTransactionDefault) WithStatusCode(code int) *GetTransactionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get transaction default response
func (o *GetTransactionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get transaction default response
func (o *GetTransactionDefault) WithConfigurationVersion(configurationVersion int64) *GetTransactionDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get transaction default response
func (o *GetTransactionDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get transaction default response
func (o *GetTransactionDefault) WithPayload(payload *models.Error) *GetTransactionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get transaction default response
func (o *GetTransactionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTransactionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
