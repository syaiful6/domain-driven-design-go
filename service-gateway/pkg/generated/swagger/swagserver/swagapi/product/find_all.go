// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	swagmodel "github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagmodel"
)

// FindAllHandlerFunc turns a function with the right signature into a find all handler
type FindAllHandlerFunc func(FindAllParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindAllHandlerFunc) Handle(params FindAllParams) middleware.Responder {
	return fn(params)
}

// FindAllHandler interface for that can handle valid find all params
type FindAllHandler interface {
	Handle(FindAllParams) middleware.Responder
}

// NewFindAll creates a new http.Handler for the find all operation
func NewFindAll(ctx *middleware.Context, handler FindAllHandler) *FindAll {
	return &FindAll{Context: ctx, Handler: handler}
}

/*FindAll swagger:route GET /product product findAll

FindAll find all API

*/
type FindAll struct {
	Context *middleware.Context
	Handler FindAllHandler
}

func (o *FindAll) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFindAllParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// FindAllOKBody find all o k body
// swagger:model FindAllOKBody
type FindAllOKBody struct {

	// pagination
	Pagination *swagmodel.Pagination `json:"pagination,omitempty"`

	// rows
	Rows []*swagmodel.Product `json:"rows"`
}

// Validate validates this find all o k body
func (o *FindAllOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRows(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindAllOKBody) validatePagination(formats strfmt.Registry) error {

	if swag.IsZero(o.Pagination) { // not required
		return nil
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("findAllOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

func (o *FindAllOKBody) validateRows(formats strfmt.Registry) error {

	if swag.IsZero(o.Rows) { // not required
		return nil
	}

	for i := 0; i < len(o.Rows); i++ {
		if swag.IsZero(o.Rows[i]) { // not required
			continue
		}

		if o.Rows[i] != nil {
			if err := o.Rows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("findAllOK" + "." + "rows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *FindAllOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FindAllOKBody) UnmarshalBinary(b []byte) error {
	var res FindAllOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
