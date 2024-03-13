// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResultInput result input
//
// swagger:model resultInput
type ResultInput struct {

	// score
	// Required: true
	Score *ScoreEnum `json:"score"`
}

// Validate validates this result input
func (m *ResultInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScore(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResultInput) validateScore(formats strfmt.Registry) error {

	if err := validate.Required("score", "body", m.Score); err != nil {
		return err
	}

	if err := validate.Required("score", "body", m.Score); err != nil {
		return err
	}

	if m.Score != nil {
		if err := m.Score.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("score")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this result input based on the context it is used
func (m *ResultInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScore(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResultInput) contextValidateScore(ctx context.Context, formats strfmt.Registry) error {

	if m.Score != nil {
		if err := m.Score.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("score")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResultInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResultInput) UnmarshalBinary(b []byte) error {
	var res ResultInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}