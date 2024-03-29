// Code generated by go-swagger; DO NOT EDIT.

package results

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/jle02/ShyftLabs-Takehome/gen/models"
)

// GetResultsScoreOKCode is the HTTP code returned for type GetResultsScoreOK
const GetResultsScoreOKCode int = 200

/*GetResultsScoreOK list the results scores

swagger:response getResultsScoreOK
*/
type GetResultsScoreOK struct {

	/*
	  In: Body
	*/
	Payload []models.ScoreEnum `json:"body,omitempty"`
}

// NewGetResultsScoreOK creates GetResultsScoreOK with default headers values
func NewGetResultsScoreOK() *GetResultsScoreOK {

	return &GetResultsScoreOK{}
}

// WithPayload adds the payload to the get results score o k response
func (o *GetResultsScoreOK) WithPayload(payload []models.ScoreEnum) *GetResultsScoreOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get results score o k response
func (o *GetResultsScoreOK) SetPayload(payload []models.ScoreEnum) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetResultsScoreOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]models.ScoreEnum, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetResultsScoreDefault generic error response

swagger:response getResultsScoreDefault
*/
type GetResultsScoreDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetResultsScoreDefault creates GetResultsScoreDefault with default headers values
func NewGetResultsScoreDefault(code int) *GetResultsScoreDefault {
	if code <= 0 {
		code = 500
	}

	return &GetResultsScoreDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get results score default response
func (o *GetResultsScoreDefault) WithStatusCode(code int) *GetResultsScoreDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get results score default response
func (o *GetResultsScoreDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get results score default response
func (o *GetResultsScoreDefault) WithPayload(payload *models.Error) *GetResultsScoreDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get results score default response
func (o *GetResultsScoreDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetResultsScoreDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
