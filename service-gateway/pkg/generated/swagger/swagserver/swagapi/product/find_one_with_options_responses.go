// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	swagmodel "github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagmodel"
)

// FindOneWithOptionsOKCode is the HTTP code returned for type FindOneWithOptionsOK
const FindOneWithOptionsOKCode int = 200

/*FindOneWithOptionsOK OK

swagger:response findOneWithOptionsOK
*/
type FindOneWithOptionsOK struct {

	/*
	  In: Body
	*/
	Payload *FindOneWithOptionsOKBody `json:"body,omitempty"`
}

// NewFindOneWithOptionsOK creates FindOneWithOptionsOK with default headers values
func NewFindOneWithOptionsOK() *FindOneWithOptionsOK {

	return &FindOneWithOptionsOK{}
}

// WithPayload adds the payload to the find one with options o k response
func (o *FindOneWithOptionsOK) WithPayload(payload *FindOneWithOptionsOKBody) *FindOneWithOptionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find one with options o k response
func (o *FindOneWithOptionsOK) SetPayload(payload *FindOneWithOptionsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindOneWithOptionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*FindOneWithOptionsDefault error

swagger:response findOneWithOptionsDefault
*/
type FindOneWithOptionsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *swagmodel.Exception `json:"body,omitempty"`
}

// NewFindOneWithOptionsDefault creates FindOneWithOptionsDefault with default headers values
func NewFindOneWithOptionsDefault(code int) *FindOneWithOptionsDefault {
	if code <= 0 {
		code = 500
	}

	return &FindOneWithOptionsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find one with options default response
func (o *FindOneWithOptionsDefault) WithStatusCode(code int) *FindOneWithOptionsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find one with options default response
func (o *FindOneWithOptionsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find one with options default response
func (o *FindOneWithOptionsDefault) WithPayload(payload *swagmodel.Exception) *FindOneWithOptionsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find one with options default response
func (o *FindOneWithOptionsDefault) SetPayload(payload *swagmodel.Exception) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindOneWithOptionsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
