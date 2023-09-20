/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// UserAPIController binds http requests to an api service and writes the service results to the http response
type UserAPIController struct {
	service UserAPIServicer
	errorHandler ErrorHandler
}

// UserAPIOption for how the controller is set up.
type UserAPIOption func(*UserAPIController)

// WithUserAPIErrorHandler inject ErrorHandler into controller
func WithUserAPIErrorHandler(h ErrorHandler) UserAPIOption {
	return func(c *UserAPIController) {
		c.errorHandler = h
	}
}

// NewUserAPIController creates a default api controller
func NewUserAPIController(s UserAPIServicer, opts ...UserAPIOption) Router {
	controller := &UserAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the UserAPIController
func (c *UserAPIController) Routes() Routes {
	return Routes{
		"CreateUser": Route{
			strings.ToUpper("Post"),
			"/v2/user",
			c.CreateUser,
		},
		"CreateUsersWithArrayInput": Route{
			strings.ToUpper("Post"),
			"/v2/user/createWithArray",
			c.CreateUsersWithArrayInput,
		},
		"CreateUsersWithListInput": Route{
			strings.ToUpper("Post"),
			"/v2/user/createWithList",
			c.CreateUsersWithListInput,
		},
		"DeleteUser": Route{
			strings.ToUpper("Delete"),
			"/v2/user/{username}",
			c.DeleteUser,
		},
		"GetUserByName": Route{
			strings.ToUpper("Get"),
			"/v2/user/{username}",
			c.GetUserByName,
		},
		"LoginUser": Route{
			strings.ToUpper("Get"),
			"/v2/user/login",
			c.LoginUser,
		},
		"LogoutUser": Route{
			strings.ToUpper("Get"),
			"/v2/user/logout",
			c.LogoutUser,
		},
		"UpdateUser": Route{
			strings.ToUpper("Put"),
			"/v2/user/{username}",
			c.UpdateUser,
		},
	}
}

// CreateUser - Create user
func (c *UserAPIController) CreateUser(w http.ResponseWriter, r *http.Request) {
	userParam := User{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&userParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserRequired(userParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUserConstraints(userParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateUser(r.Context(), userParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateUsersWithArrayInput - Creates list of users with given input array
func (c *UserAPIController) CreateUsersWithArrayInput(w http.ResponseWriter, r *http.Request) {
	userParam := []User{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&userParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	for _, el := range userParam {
		if err := AssertUserRequired(el); err != nil {
			c.errorHandler(w, r, err, nil)
			return
		}
	}
	result, err := c.service.CreateUsersWithArrayInput(r.Context(), userParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateUsersWithListInput - Creates list of users with given input array
func (c *UserAPIController) CreateUsersWithListInput(w http.ResponseWriter, r *http.Request) {
	userParam := []User{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&userParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	for _, el := range userParam {
		if err := AssertUserRequired(el); err != nil {
			c.errorHandler(w, r, err, nil)
			return
		}
	}
	result, err := c.service.CreateUsersWithListInput(r.Context(), userParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteUser - Delete user
func (c *UserAPIController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usernameParam := params["username"]
	result, err := c.service.DeleteUser(r.Context(), usernameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetUserByName - Get user by user name
func (c *UserAPIController) GetUserByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usernameParam := params["username"]
	result, err := c.service.GetUserByName(r.Context(), usernameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// LoginUser - Logs user into the system
func (c *UserAPIController) LoginUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	usernameParam := query.Get("username")
	passwordParam := query.Get("password")
	result, err := c.service.LoginUser(r.Context(), usernameParam, passwordParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// LogoutUser - Logs out current logged in user session
func (c *UserAPIController) LogoutUser(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.LogoutUser(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateUser - Updated user
func (c *UserAPIController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usernameParam := params["username"]
	userParam := User{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&userParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserRequired(userParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUserConstraints(userParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateUser(r.Context(), usernameParam, userParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
