// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindAllParams creates a new FindAllParams object
// with the default values initialized.
func NewFindAllParams() FindAllParams {

	var (
		// initialize parameters with default values

		currentPageOffsetDefault = int32(0)
		itemCountPerPageDefault  = int32(10)
	)

	return FindAllParams{
		CurrentPageOffset: &currentPageOffsetDefault,

		ItemCountPerPage: &itemCountPerPageDefault,
	}
}

// FindAllParams contains all the bound params for the find all operation
// typically these are obtained from a http.Request
//
// swagger:parameters findAll
type FindAllParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	  Default: 0
	*/
	CurrentPageOffset *int32
	/*
	  In: query
	  Default: 10
	*/
	ItemCountPerPage *int32
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewFindAllParams() beforehand.
func (o *FindAllParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCurrentPageOffset, qhkCurrentPageOffset, _ := qs.GetOK("currentPageOffset")
	if err := o.bindCurrentPageOffset(qCurrentPageOffset, qhkCurrentPageOffset, route.Formats); err != nil {
		res = append(res, err)
	}

	qItemCountPerPage, qhkItemCountPerPage, _ := qs.GetOK("itemCountPerPage")
	if err := o.bindItemCountPerPage(qItemCountPerPage, qhkItemCountPerPage, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindAllParams) bindCurrentPageOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewFindAllParams()
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("currentPageOffset", "query", "int32", raw)
	}
	o.CurrentPageOffset = &value

	return nil
}

func (o *FindAllParams) bindItemCountPerPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewFindAllParams()
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("itemCountPerPage", "query", "int32", raw)
	}
	o.ItemCountPerPage = &value

	return nil
}
