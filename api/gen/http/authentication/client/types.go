// Code generated by goa v3.16.2, DO NOT EDIT.
//
// Authentication HTTP client types
//
// Command:
// $ goa gen mai.today/api/design

package client

import (
	goa "goa.design/goa/v3/pkg"
	authentication "mai.today/api/gen/authentication"
)

// SignInRequestBody is the type of the "Authentication" service "SignIn"
// endpoint HTTP request body.
type SignInRequestBody struct {
	// Firebase Id Token
	FirebaseIDToken string `form:"firebaseIdToken" json:"firebaseIdToken" xml:"firebaseIdToken"`
}

// SignInTokenErrorResponseBody is the type of the "Authentication" service
// "SignIn" endpoint HTTP response body for the "token_error" error.
type SignInTokenErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// NewSignInRequestBody builds the HTTP request body from the payload of the
// "SignIn" endpoint of the "Authentication" service.
func NewSignInRequestBody(p *authentication.SignInPayload) *SignInRequestBody {
	body := &SignInRequestBody{
		FirebaseIDToken: p.FirebaseIDToken,
	}
	return body
}

// NewSignInTokenError builds a Authentication service SignIn endpoint
// token_error error.
func NewSignInTokenError(body *SignInTokenErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateSignInTokenErrorResponseBody runs the validations defined on
// SignIn_token_error_Response_Body
func ValidateSignInTokenErrorResponseBody(body *SignInTokenErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}
