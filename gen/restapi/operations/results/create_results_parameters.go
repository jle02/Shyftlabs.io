// Code generated by go-swagger; DO NOT EDIT.

package results

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/jle02/ShyftLabs-Takehome/gen/models"
)

// NewCreateResultsParams creates a new CreateResultsParams object
//
// There are no default values defined in the spec.
func NewCreateResultsParams() CreateResultsParams {

	return CreateResultsParams{}
}

// CreateResultsParams contains all the bound params for the create results operation
// typically these are obtained from a http.Request
//
// swagger:parameters createResults
type CreateResultsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Body *models.ResultInput
	/*
	  Required: true
	  In: path
	*/
	CourseID int64
	/*
	  Required: true
	  In: path
	*/
	StudentID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateResultsParams() beforehand.
func (o *CreateResultsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ResultInput
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	}

	rCourseID, rhkCourseID, _ := route.Params.GetOK("courseId")
	if err := o.bindCourseID(rCourseID, rhkCourseID, route.Formats); err != nil {
		res = append(res, err)
	}

	rStudentID, rhkStudentID, _ := route.Params.GetOK("studentId")
	if err := o.bindStudentID(rStudentID, rhkStudentID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCourseID binds and validates parameter CourseID from path.
func (o *CreateResultsParams) bindCourseID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("courseId", "path", "int64", raw)
	}
	o.CourseID = value

	return nil
}

// bindStudentID binds and validates parameter StudentID from path.
func (o *CreateResultsParams) bindStudentID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("studentId", "path", "int64", raw)
	}
	o.StudentID = value

	return nil
}