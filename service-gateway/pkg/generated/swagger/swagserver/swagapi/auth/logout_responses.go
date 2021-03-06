// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	swagmodel "github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagmodel"
)

// LogoutOKCode is the HTTP code returned for type LogoutOK
const LogoutOKCode int = 200

/*LogoutOK OK

swagger:response logoutOK
*/
type LogoutOK struct {

	/*
	  In: Body
	*/
	Payload swagmodel.Empty `json:"body,omitempty"`
}

// NewLogoutOK creates LogoutOK with default headers values
func NewLogoutOK() *LogoutOK {

	return &LogoutOK{}
}

// WithPayload adds the payload to the logout o k response
func (o *LogoutOK) WithPayload(payload swagmodel.Empty) *LogoutOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the logout o k response
func (o *LogoutOK) SetPayload(payload swagmodel.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LogoutOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*LogoutDefault error

swagger:response logoutDefault
*/
type LogoutDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *swagmodel.Exception `json:"body,omitempty"`
}

// NewLogoutDefault creates LogoutDefault with default headers values
func NewLogoutDefault(code int) *LogoutDefault {
	if code <= 0 {
		code = 500
	}

	return &LogoutDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the logout default response
func (o *LogoutDefault) WithStatusCode(code int) *LogoutDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the logout default response
func (o *LogoutDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the logout default response
func (o *LogoutDefault) WithPayload(payload *swagmodel.Exception) *LogoutDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the logout default response
func (o *LogoutDefault) SetPayload(payload *swagmodel.Exception) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LogoutDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
