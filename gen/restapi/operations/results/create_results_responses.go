// Code generated by go-swagger; DO NOT EDIT.

package results

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/jle02/ShyftLabs-Takehome/gen/models"
)

// CreateResultsCreatedCode is the HTTP code returned for type CreateResultsCreated
const CreateResultsCreatedCode int = 201

/*CreateResultsCreated create a result

swagger:response createResultsCreated
*/
type CreateResultsCreated struct {

	/*
	  In: Body
	*/
	Payload *models.ResultOutput `json:"body,omitempty"`
}

// NewCreateResultsCreated creates CreateResultsCreated with default headers values
func NewCreateResultsCreated() *CreateResultsCreated {

	return &CreateResultsCreated{}
}

// WithPayload adds the payload to the create results created response
func (o *CreateResultsCreated) WithPayload(payload *models.ResultOutput) *CreateResultsCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create results created response
func (o *CreateResultsCreated) SetPayload(payload *models.ResultOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateResultsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateResultsDefault generic error response

swagger:response createResultsDefault
*/
type CreateResultsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateResultsDefault creates CreateResultsDefault with default headers values
func NewCreateResultsDefault(code int) *CreateResultsDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateResultsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create results default response
func (o *CreateResultsDefault) WithStatusCode(code int) *CreateResultsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create results default response
func (o *CreateResultsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create results default response
func (o *CreateResultsDefault) WithPayload(payload *models.Error) *CreateResultsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create results default response
func (o *CreateResultsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateResultsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
