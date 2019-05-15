/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */ // Code generated by go-swagger; DO NOT EDIT.

package contextionary_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateC11yWordsOKCode is the HTTP code returned for type WeaviateC11yWordsOK
const WeaviateC11yWordsOKCode int = 200

/*WeaviateC11yWordsOK Successful response.

swagger:response weaviateC11yWordsOK
*/
type WeaviateC11yWordsOK struct {

	/*
	  In: Body
	*/
	Payload *models.C11yWordsResponse `json:"body,omitempty"`
}

// NewWeaviateC11yWordsOK creates WeaviateC11yWordsOK with default headers values
func NewWeaviateC11yWordsOK() *WeaviateC11yWordsOK {

	return &WeaviateC11yWordsOK{}
}

// WithPayload adds the payload to the weaviate c11y words o k response
func (o *WeaviateC11yWordsOK) WithPayload(payload *models.C11yWordsResponse) *WeaviateC11yWordsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate c11y words o k response
func (o *WeaviateC11yWordsOK) SetPayload(payload *models.C11yWordsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateC11yWordsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateC11yWordsBadRequestCode is the HTTP code returned for type WeaviateC11yWordsBadRequest
const WeaviateC11yWordsBadRequestCode int = 400

/*WeaviateC11yWordsBadRequest Incorrect request

swagger:response weaviateC11yWordsBadRequest
*/
type WeaviateC11yWordsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateC11yWordsBadRequest creates WeaviateC11yWordsBadRequest with default headers values
func NewWeaviateC11yWordsBadRequest() *WeaviateC11yWordsBadRequest {

	return &WeaviateC11yWordsBadRequest{}
}

// WithPayload adds the payload to the weaviate c11y words bad request response
func (o *WeaviateC11yWordsBadRequest) WithPayload(payload *models.ErrorResponse) *WeaviateC11yWordsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate c11y words bad request response
func (o *WeaviateC11yWordsBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateC11yWordsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateC11yWordsUnauthorizedCode is the HTTP code returned for type WeaviateC11yWordsUnauthorized
const WeaviateC11yWordsUnauthorizedCode int = 401

/*WeaviateC11yWordsUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateC11yWordsUnauthorized
*/
type WeaviateC11yWordsUnauthorized struct {
}

// NewWeaviateC11yWordsUnauthorized creates WeaviateC11yWordsUnauthorized with default headers values
func NewWeaviateC11yWordsUnauthorized() *WeaviateC11yWordsUnauthorized {

	return &WeaviateC11yWordsUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateC11yWordsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateC11yWordsForbiddenCode is the HTTP code returned for type WeaviateC11yWordsForbidden
const WeaviateC11yWordsForbiddenCode int = 403

/*WeaviateC11yWordsForbidden Insufficient permissions.

swagger:response weaviateC11yWordsForbidden
*/
type WeaviateC11yWordsForbidden struct {
}

// NewWeaviateC11yWordsForbidden creates WeaviateC11yWordsForbidden with default headers values
func NewWeaviateC11yWordsForbidden() *WeaviateC11yWordsForbidden {

	return &WeaviateC11yWordsForbidden{}
}

// WriteResponse to the client
func (o *WeaviateC11yWordsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// WeaviateC11yWordsInternalServerErrorCode is the HTTP code returned for type WeaviateC11yWordsInternalServerError
const WeaviateC11yWordsInternalServerErrorCode int = 500

/*WeaviateC11yWordsInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response weaviateC11yWordsInternalServerError
*/
type WeaviateC11yWordsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateC11yWordsInternalServerError creates WeaviateC11yWordsInternalServerError with default headers values
func NewWeaviateC11yWordsInternalServerError() *WeaviateC11yWordsInternalServerError {

	return &WeaviateC11yWordsInternalServerError{}
}

// WithPayload adds the payload to the weaviate c11y words internal server error response
func (o *WeaviateC11yWordsInternalServerError) WithPayload(payload *models.ErrorResponse) *WeaviateC11yWordsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate c11y words internal server error response
func (o *WeaviateC11yWordsInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateC11yWordsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateC11yWordsNotImplementedCode is the HTTP code returned for type WeaviateC11yWordsNotImplemented
const WeaviateC11yWordsNotImplementedCode int = 501

/*WeaviateC11yWordsNotImplemented Not (yet) implemented.

swagger:response weaviateC11yWordsNotImplemented
*/
type WeaviateC11yWordsNotImplemented struct {
}

// NewWeaviateC11yWordsNotImplemented creates WeaviateC11yWordsNotImplemented with default headers values
func NewWeaviateC11yWordsNotImplemented() *WeaviateC11yWordsNotImplemented {

	return &WeaviateC11yWordsNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateC11yWordsNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(501)
}
