/*
 * XXX API
 *
 * XXXについてのAPI
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// LoginApiController binds http requests to an api service and writes the service results to the http response
type LoginApiController struct {
	service LoginApiServicer
	errorHandler ErrorHandler
}

// LoginApiOption for how the controller is set up.
type LoginApiOption func(*LoginApiController)

// WithLoginApiErrorHandler inject ErrorHandler into controller
func WithLoginApiErrorHandler(h ErrorHandler) LoginApiOption {
	return func(c *LoginApiController) {
		c.errorHandler = h
	}
}

// NewLoginApiController creates a default api controller
func NewLoginApiController(s LoginApiServicer, opts ...LoginApiOption) Router {
	controller := &LoginApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the LoginApiController
func (c *LoginApiController) Routes() Routes {
	return Routes{ 
		{
			"StartLogin",
			strings.ToUpper("Get"),
			"/login",
			c.StartLogin,
		},
	}
}

// StartLogin - Login
func (c *LoginApiController) StartLogin(w http.ResponseWriter, r *http.Request) {
	loginBodyParam := LoginBody{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&loginBodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertLoginBodyRequired(loginBodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.StartLogin(r.Context(), loginBodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}