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

// DeleteLogTargetPeerAcceptedCode is the HTTP code returned for type DeleteLogTargetPeerAccepted
const DeleteLogTargetPeerAcceptedCode int = 202

/*
DeleteLogTargetPeerAccepted Configuration change accepted and reload requested

swagger:response deleteLogTargetPeerAccepted
*/
type DeleteLogTargetPeerAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteLogTargetPeerAccepted creates DeleteLogTargetPeerAccepted with default headers values
func NewDeleteLogTargetPeerAccepted() *DeleteLogTargetPeerAccepted {

	return &DeleteLogTargetPeerAccepted{}
}

// WithReloadID adds the reloadId to the delete log target peer accepted response
func (o *DeleteLogTargetPeerAccepted) WithReloadID(reloadID string) *DeleteLogTargetPeerAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete log target peer accepted response
func (o *DeleteLogTargetPeerAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteLogTargetPeerAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteLogTargetPeerNoContentCode is the HTTP code returned for type DeleteLogTargetPeerNoContent
const DeleteLogTargetPeerNoContentCode int = 204

/*
DeleteLogTargetPeerNoContent Log Target deleted

swagger:response deleteLogTargetPeerNoContent
*/
type DeleteLogTargetPeerNoContent struct {
}

// NewDeleteLogTargetPeerNoContent creates DeleteLogTargetPeerNoContent with default headers values
func NewDeleteLogTargetPeerNoContent() *DeleteLogTargetPeerNoContent {

	return &DeleteLogTargetPeerNoContent{}
}

// WriteResponse to the client
func (o *DeleteLogTargetPeerNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteLogTargetPeerNotFoundCode is the HTTP code returned for type DeleteLogTargetPeerNotFound
const DeleteLogTargetPeerNotFoundCode int = 404

/*
DeleteLogTargetPeerNotFound The specified resource was not found

swagger:response deleteLogTargetPeerNotFound
*/
type DeleteLogTargetPeerNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteLogTargetPeerNotFound creates DeleteLogTargetPeerNotFound with default headers values
func NewDeleteLogTargetPeerNotFound() *DeleteLogTargetPeerNotFound {

	return &DeleteLogTargetPeerNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete log target peer not found response
func (o *DeleteLogTargetPeerNotFound) WithConfigurationVersion(configurationVersion string) *DeleteLogTargetPeerNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete log target peer not found response
func (o *DeleteLogTargetPeerNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete log target peer not found response
func (o *DeleteLogTargetPeerNotFound) WithPayload(payload *models.Error) *DeleteLogTargetPeerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete log target peer not found response
func (o *DeleteLogTargetPeerNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteLogTargetPeerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
DeleteLogTargetPeerDefault General Error

swagger:response deleteLogTargetPeerDefault
*/
type DeleteLogTargetPeerDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteLogTargetPeerDefault creates DeleteLogTargetPeerDefault with default headers values
func NewDeleteLogTargetPeerDefault(code int) *DeleteLogTargetPeerDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteLogTargetPeerDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete log target peer default response
func (o *DeleteLogTargetPeerDefault) WithStatusCode(code int) *DeleteLogTargetPeerDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete log target peer default response
func (o *DeleteLogTargetPeerDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete log target peer default response
func (o *DeleteLogTargetPeerDefault) WithConfigurationVersion(configurationVersion string) *DeleteLogTargetPeerDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete log target peer default response
func (o *DeleteLogTargetPeerDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete log target peer default response
func (o *DeleteLogTargetPeerDefault) WithPayload(payload *models.Error) *DeleteLogTargetPeerDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete log target peer default response
func (o *DeleteLogTargetPeerDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteLogTargetPeerDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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