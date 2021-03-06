// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/home-IoT/api-samsungtv/gen/models"
)

// PostKeyAcceptedCode is the HTTP code returned for type PostKeyAccepted
const PostKeyAcceptedCode int = 202

/*PostKeyAccepted Accepted

swagger:response postKeyAccepted
*/
type PostKeyAccepted struct {
}

// NewPostKeyAccepted creates PostKeyAccepted with default headers values
func NewPostKeyAccepted() *PostKeyAccepted {

	return &PostKeyAccepted{}
}

// WriteResponse to the client
func (o *PostKeyAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

/*PostKeyDefault Error

swagger:response postKeyDefault
*/
type PostKeyDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload models.ErrorResponse `json:"body,omitempty"`
}

// NewPostKeyDefault creates PostKeyDefault with default headers values
func NewPostKeyDefault(code int) *PostKeyDefault {
	if code <= 0 {
		code = 500
	}

	return &PostKeyDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post key default response
func (o *PostKeyDefault) WithStatusCode(code int) *PostKeyDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post key default response
func (o *PostKeyDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post key default response
func (o *PostKeyDefault) WithPayload(payload models.ErrorResponse) *PostKeyDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post key default response
func (o *PostKeyDefault) SetPayload(payload models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostKeyDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
