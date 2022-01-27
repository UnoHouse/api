// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/UnoHouse/api/swagger/models"
)

// GetHealthcheckOKCode is the HTTP code returned for type GetHealthcheckOK
const GetHealthcheckOKCode int = 200

/*GetHealthcheckOK OK

swagger:response getHealthcheckOK
*/
type GetHealthcheckOK struct {
}

// NewGetHealthcheckOK creates GetHealthcheckOK with default headers values
func NewGetHealthcheckOK() *GetHealthcheckOK {

	return &GetHealthcheckOK{}
}

// WriteResponse to the client
func (o *GetHealthcheckOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetHealthcheckBadRequestCode is the HTTP code returned for type GetHealthcheckBadRequest
const GetHealthcheckBadRequestCode int = 400

/*GetHealthcheckBadRequest Bad request

swagger:response getHealthcheckBadRequest
*/
type GetHealthcheckBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetHealthcheckBadRequest creates GetHealthcheckBadRequest with default headers values
func NewGetHealthcheckBadRequest() *GetHealthcheckBadRequest {

	return &GetHealthcheckBadRequest{}
}

// WithPayload adds the payload to the get healthcheck bad request response
func (o *GetHealthcheckBadRequest) WithPayload(payload *models.ErrorResponse) *GetHealthcheckBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get healthcheck bad request response
func (o *GetHealthcheckBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHealthcheckBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHealthcheckInternalServerErrorCode is the HTTP code returned for type GetHealthcheckInternalServerError
const GetHealthcheckInternalServerErrorCode int = 500

/*GetHealthcheckInternalServerError Internal server error

swagger:response getHealthcheckInternalServerError
*/
type GetHealthcheckInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetHealthcheckInternalServerError creates GetHealthcheckInternalServerError with default headers values
func NewGetHealthcheckInternalServerError() *GetHealthcheckInternalServerError {

	return &GetHealthcheckInternalServerError{}
}

// WithPayload adds the payload to the get healthcheck internal server error response
func (o *GetHealthcheckInternalServerError) WithPayload(payload *models.ErrorResponse) *GetHealthcheckInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get healthcheck internal server error response
func (o *GetHealthcheckInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHealthcheckInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}