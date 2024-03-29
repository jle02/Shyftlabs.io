// Code generated by go-swagger; DO NOT EDIT.

package students

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateStudentsHandlerFunc turns a function with the right signature into a create students handler
type CreateStudentsHandlerFunc func(CreateStudentsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateStudentsHandlerFunc) Handle(params CreateStudentsParams) middleware.Responder {
	return fn(params)
}

// CreateStudentsHandler interface for that can handle valid create students params
type CreateStudentsHandler interface {
	Handle(CreateStudentsParams) middleware.Responder
}

// NewCreateStudents creates a new http.Handler for the create students operation
func NewCreateStudents(ctx *middleware.Context, handler CreateStudentsHandler) *CreateStudents {
	return &CreateStudents{Context: ctx, Handler: handler}
}

/* CreateStudents swagger:route POST /students students createStudents

CreateStudents create students API

*/
type CreateStudents struct {
	Context *middleware.Context
	Handler CreateStudentsHandler
}

func (o *CreateStudents) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateStudentsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
