// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1FindIpsRequest an ip address that can be attached to a machine
// swagger:model v1.FindIPsRequest
type V1FindIpsRequest struct {

	// the address (ipv4 or ipv6) of this ip
	// Required: true
	Ipaddress *string `json:"ipaddress"`

	// the machine this ip address belongs to, empty if not strong coupled
	// Required: true
	Machineid *string `json:"machineid"`

	// the network this ip allocate request address belongs to
	// Required: true
	Networkid *string `json:"networkid"`

	// the prefix of the network this ip address belongs to
	// Required: true
	Networkprefix *string `json:"networkprefix"`

	// the project this ip address belongs to
	// Required: true
	Projectid *string `json:"projectid"`
}

// Validate validates this v1 find ips request
func (m *V1FindIpsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIpaddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkprefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1FindIpsRequest) validateIpaddress(formats strfmt.Registry) error {

	if err := validate.Required("ipaddress", "body", m.Ipaddress); err != nil {
		return err
	}

	return nil
}

func (m *V1FindIpsRequest) validateMachineid(formats strfmt.Registry) error {

	if err := validate.Required("machineid", "body", m.Machineid); err != nil {
		return err
	}

	return nil
}

func (m *V1FindIpsRequest) validateNetworkid(formats strfmt.Registry) error {

	if err := validate.Required("networkid", "body", m.Networkid); err != nil {
		return err
	}

	return nil
}

func (m *V1FindIpsRequest) validateNetworkprefix(formats strfmt.Registry) error {

	if err := validate.Required("networkprefix", "body", m.Networkprefix); err != nil {
		return err
	}

	return nil
}

func (m *V1FindIpsRequest) validateProjectid(formats strfmt.Registry) error {

	if err := validate.Required("projectid", "body", m.Projectid); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1FindIpsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1FindIpsRequest) UnmarshalBinary(b []byte) error {
	var res V1FindIpsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}