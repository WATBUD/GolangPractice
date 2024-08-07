// Code generated by goa v3.16.2, DO NOT EDIT.
//
// Authentication client HTTP transport
//
// Command:
// $ goa gen mai.today/api/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the Authentication service endpoint HTTP clients.
type Client struct {
	// SignIn Doer is the HTTP client used to make requests to the SignIn endpoint.
	SignInDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the Authentication service
// servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		SignInDoer:          doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// SignIn returns an endpoint that makes HTTP requests to the Authentication
// service SignIn server.
func (c *Client) SignIn() goa.Endpoint {
	var (
		encodeRequest  = EncodeSignInRequest(c.encoder)
		decodeResponse = DecodeSignInResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildSignInRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SignInDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("Authentication", "SignIn", err)
		}
		return decodeResponse(resp)
	}
}
