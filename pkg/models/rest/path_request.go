// Code generated by go-swagger; DO NOT EDIT.

package rest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PathRequest path request
//
// swagger:model PathRequest
type PathRequest struct {

	// flights
	// Required: true
	Flights []*Flight `json:"flights"`
}

// Validate validates this path request
func (m *PathRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlights(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PathRequest) validateFlights(formats strfmt.Registry) error {

	if err := validate.Required("flights", "body", m.Flights); err != nil {
		return err
	}

	for i := 0; i < len(m.Flights); i++ {
		if swag.IsZero(m.Flights[i]) { // not required
			continue
		}

		if m.Flights[i] != nil {
			if err := m.Flights[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("flights" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("flights" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this path request based on the context it is used
func (m *PathRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFlights(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PathRequest) contextValidateFlights(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Flights); i++ {

		if m.Flights[i] != nil {
			if err := m.Flights[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("flights" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("flights" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PathRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PathRequest) UnmarshalBinary(b []byte) error {
	var res PathRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}