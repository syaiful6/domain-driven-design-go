// Code generated by go-swagger; DO NOT EDIT.

package swagmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetCartItemsOKBody get cart items o k body
// swagger:model getCartItemsOKBody
type GetCartItemsOKBody struct {

	// cart
	Cart *Cart `json:"cart,omitempty"`

	// cart item list
	CartItemList []*GetCartItemsOKBodyCartItemListItems `json:"cartItemList"`
}

// Validate validates this get cart items o k body
func (m *GetCartItemsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCartItemList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCartItemsOKBody) validateCart(formats strfmt.Registry) error {

	if swag.IsZero(m.Cart) { // not required
		return nil
	}

	if m.Cart != nil {
		if err := m.Cart.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cart")
			}
			return err
		}
	}

	return nil
}

func (m *GetCartItemsOKBody) validateCartItemList(formats strfmt.Registry) error {

	if swag.IsZero(m.CartItemList) { // not required
		return nil
	}

	for i := 0; i < len(m.CartItemList); i++ {
		if swag.IsZero(m.CartItemList[i]) { // not required
			continue
		}

		if m.CartItemList[i] != nil {
			if err := m.CartItemList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cartItemList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCartItemsOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCartItemsOKBody) UnmarshalBinary(b []byte) error {
	var res GetCartItemsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
