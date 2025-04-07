// Code generated by go-swagger; DO NOT EDIT.

package ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/DropsWeb/todogo/pkg/swagger/models"
)

// PostAuthLoginOKCode is the HTTP code returned for type PostAuthLoginOK
const PostAuthLoginOKCode int = 200

/*
PostAuthLoginOK Successful login

swagger:response postAuthLoginOK
*/
type PostAuthLoginOK struct {

	/*
	  In: Body
	*/
	Payload *models.LoginSuccess `json:"body,omitempty"`
}

// NewPostAuthLoginOK creates PostAuthLoginOK with default headers values
func NewPostAuthLoginOK() *PostAuthLoginOK {

	return &PostAuthLoginOK{}
}

// WithPayload adds the payload to the post auth login o k response
func (o *PostAuthLoginOK) WithPayload(payload *models.LoginSuccess) *PostAuthLoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth login o k response
func (o *PostAuthLoginOK) SetPayload(payload *models.LoginSuccess) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthLoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAuthLoginBadRequestCode is the HTTP code returned for type PostAuthLoginBadRequest
const PostAuthLoginBadRequestCode int = 400

/*
PostAuthLoginBadRequest Bad Request

swagger:response postAuthLoginBadRequest
*/
type PostAuthLoginBadRequest struct {
}

// NewPostAuthLoginBadRequest creates PostAuthLoginBadRequest with default headers values
func NewPostAuthLoginBadRequest() *PostAuthLoginBadRequest {

	return &PostAuthLoginBadRequest{}
}

// WriteResponse to the client
func (o *PostAuthLoginBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PostAuthLoginUnauthorizedCode is the HTTP code returned for type PostAuthLoginUnauthorized
const PostAuthLoginUnauthorizedCode int = 401

/*
PostAuthLoginUnauthorized Unauthorized

swagger:response postAuthLoginUnauthorized
*/
type PostAuthLoginUnauthorized struct {
}

// NewPostAuthLoginUnauthorized creates PostAuthLoginUnauthorized with default headers values
func NewPostAuthLoginUnauthorized() *PostAuthLoginUnauthorized {

	return &PostAuthLoginUnauthorized{}
}

// WriteResponse to the client
func (o *PostAuthLoginUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// PostAuthLoginNotFoundCode is the HTTP code returned for type PostAuthLoginNotFound
const PostAuthLoginNotFoundCode int = 404

/*
PostAuthLoginNotFound User not found

swagger:response postAuthLoginNotFound
*/
type PostAuthLoginNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostAuthLoginNotFound creates PostAuthLoginNotFound with default headers values
func NewPostAuthLoginNotFound() *PostAuthLoginNotFound {

	return &PostAuthLoginNotFound{}
}

// WithPayload adds the payload to the post auth login not found response
func (o *PostAuthLoginNotFound) WithPayload(payload string) *PostAuthLoginNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth login not found response
func (o *PostAuthLoginNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthLoginNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostAuthLoginInternalServerErrorCode is the HTTP code returned for type PostAuthLoginInternalServerError
const PostAuthLoginInternalServerErrorCode int = 500

/*
PostAuthLoginInternalServerError Server error

swagger:response postAuthLoginInternalServerError
*/
type PostAuthLoginInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostAuthLoginInternalServerError creates PostAuthLoginInternalServerError with default headers values
func NewPostAuthLoginInternalServerError() *PostAuthLoginInternalServerError {

	return &PostAuthLoginInternalServerError{}
}

// WithPayload adds the payload to the post auth login internal server error response
func (o *PostAuthLoginInternalServerError) WithPayload(payload string) *PostAuthLoginInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth login internal server error response
func (o *PostAuthLoginInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthLoginInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
