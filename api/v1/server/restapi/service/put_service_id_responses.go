// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// PutServiceIDOKCode is the HTTP code returned for type PutServiceIDOK
const PutServiceIDOKCode int = 200

/*PutServiceIDOK Updated

swagger:response putServiceIdOK
*/
type PutServiceIDOK struct {
}

// NewPutServiceIDOK creates PutServiceIDOK with default headers values
func NewPutServiceIDOK() *PutServiceIDOK {

	return &PutServiceIDOK{}
}

// WriteResponse to the client
func (o *PutServiceIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PutServiceIDCreatedCode is the HTTP code returned for type PutServiceIDCreated
const PutServiceIDCreatedCode int = 201

/*PutServiceIDCreated Created

swagger:response putServiceIdCreated
*/
type PutServiceIDCreated struct {
}

// NewPutServiceIDCreated creates PutServiceIDCreated with default headers values
func NewPutServiceIDCreated() *PutServiceIDCreated {

	return &PutServiceIDCreated{}
}

// WriteResponse to the client
func (o *PutServiceIDCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PutServiceIDInvalidFrontendCode is the HTTP code returned for type PutServiceIDInvalidFrontend
const PutServiceIDInvalidFrontendCode int = 460

/*PutServiceIDInvalidFrontend Invalid frontend in service configuration

swagger:response putServiceIdInvalidFrontend
*/
type PutServiceIDInvalidFrontend struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPutServiceIDInvalidFrontend creates PutServiceIDInvalidFrontend with default headers values
func NewPutServiceIDInvalidFrontend() *PutServiceIDInvalidFrontend {

	return &PutServiceIDInvalidFrontend{}
}

// WithPayload adds the payload to the put service Id invalid frontend response
func (o *PutServiceIDInvalidFrontend) WithPayload(payload models.Error) *PutServiceIDInvalidFrontend {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put service Id invalid frontend response
func (o *PutServiceIDInvalidFrontend) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutServiceIDInvalidFrontend) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(460)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PutServiceIDInvalidBackendCode is the HTTP code returned for type PutServiceIDInvalidBackend
const PutServiceIDInvalidBackendCode int = 461

/*PutServiceIDInvalidBackend Invalid backend in service configuration

swagger:response putServiceIdInvalidBackend
*/
type PutServiceIDInvalidBackend struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPutServiceIDInvalidBackend creates PutServiceIDInvalidBackend with default headers values
func NewPutServiceIDInvalidBackend() *PutServiceIDInvalidBackend {

	return &PutServiceIDInvalidBackend{}
}

// WithPayload adds the payload to the put service Id invalid backend response
func (o *PutServiceIDInvalidBackend) WithPayload(payload models.Error) *PutServiceIDInvalidBackend {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put service Id invalid backend response
func (o *PutServiceIDInvalidBackend) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutServiceIDInvalidBackend) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(461)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PutServiceIDFailureCode is the HTTP code returned for type PutServiceIDFailure
const PutServiceIDFailureCode int = 500

/*PutServiceIDFailure Error while creating service

swagger:response putServiceIdFailure
*/
type PutServiceIDFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPutServiceIDFailure creates PutServiceIDFailure with default headers values
func NewPutServiceIDFailure() *PutServiceIDFailure {

	return &PutServiceIDFailure{}
}

// WithPayload adds the payload to the put service Id failure response
func (o *PutServiceIDFailure) WithPayload(payload models.Error) *PutServiceIDFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put service Id failure response
func (o *PutServiceIDFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutServiceIDFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
