// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	swagmodel "github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagmodel"
)

// NewLogoutParams creates a new LogoutParams object
// no default values defined in spec.
func NewLogoutParams() LogoutParams {

	return LogoutParams{}
}

// LogoutParams contains all the bound params for the logout operation
// typically these are obtained from a http.Request
//
// swagger:parameters Logout
type LogoutParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Body swagmodel.Empty
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewLogoutParams() beforehand.
func (o *LogoutParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body swagmodel.Empty
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			// no validation on generic interface
			o.Body = body
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
