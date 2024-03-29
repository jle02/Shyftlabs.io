// Code generated by go-swagger; DO NOT EDIT.

package courses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/jle02/ShyftLabs-Takehome/gen/models"
)

// GetCoursesOKCode is the HTTP code returned for type GetCoursesOK
const GetCoursesOKCode int = 200

/*GetCoursesOK list the courses

swagger:response getCoursesOK
*/
type GetCoursesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.CourseOutput `json:"body,omitempty"`
}

// NewGetCoursesOK creates GetCoursesOK with default headers values
func NewGetCoursesOK() *GetCoursesOK {

	return &GetCoursesOK{}
}

// WithPayload adds the payload to the get courses o k response
func (o *GetCoursesOK) WithPayload(payload []*models.CourseOutput) *GetCoursesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get courses o k response
func (o *GetCoursesOK) SetPayload(payload []*models.CourseOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCoursesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.CourseOutput, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetCoursesDefault generic error response

swagger:response getCoursesDefault
*/
type GetCoursesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetCoursesDefault creates GetCoursesDefault with default headers values
func NewGetCoursesDefault(code int) *GetCoursesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetCoursesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get courses default response
func (o *GetCoursesDefault) WithStatusCode(code int) *GetCoursesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get courses default response
func (o *GetCoursesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get courses default response
func (o *GetCoursesDefault) WithPayload(payload *models.Error) *GetCoursesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get courses default response
func (o *GetCoursesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCoursesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
