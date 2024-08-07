// Code generated by goa v3.16.2, DO NOT EDIT.
//
// Base HTTP client encoders and decoders
//
// Command:
// $ goa gen mai.today/api/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	goahttp "goa.design/goa/v3/http"
	base "mai.today/api/gen/base"
)

// BuildReceiveCreateBaseRequest instantiates a HTTP request object with method
// and path set to call the "Base" service "receiveCreateBase" endpoint
func (c *Client) BuildReceiveCreateBaseRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		channel string
	)
	{
		p, ok := v.(*base.ReceiveCreateBasePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("Base", "receiveCreateBase", "*base.ReceiveCreateBasePayload", v)
		}
		channel = p.Channel
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ReceiveCreateBaseBasePath(channel)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Base", "receiveCreateBase", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeReceiveCreateBaseRequest returns an encoder for requests sent to the
// Base receiveCreateBase server.
func EncodeReceiveCreateBaseRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*base.ReceiveCreateBasePayload)
		if !ok {
			return goahttp.ErrInvalidType("Base", "receiveCreateBase", "*base.ReceiveCreateBasePayload", v)
		}
		if p.JWT != nil {
			head := *p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeReceiveCreateBaseResponse returns a decoder for responses returned by
// the Base receiveCreateBase endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeReceiveCreateBaseResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusSwitchingProtocols:
			res := NewReceiveCreateBaseCreateBaseResultSwitchingProtocols()
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Base", "receiveCreateBase", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateBaseRequest instantiates a HTTP request object with method and
// path set to call the "Base" service "CreateBase" endpoint
func (c *Client) BuildCreateBaseRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateBaseBasePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Base", "CreateBase", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateBaseRequest returns an encoder for requests sent to the Base
// CreateBase server.
func EncodeCreateBaseRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*base.CreateBasePayload)
		if !ok {
			return goahttp.ErrInvalidType("Base", "CreateBase", "*base.CreateBasePayload", v)
		}
		if p.JWT != nil {
			head := *p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewCreateBaseRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("Base", "CreateBase", err)
		}
		return nil
	}
}

// DecodeCreateBaseResponse returns a decoder for responses returned by the
// Base CreateBase endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeCreateBaseResponse may return the following errors:
//   - "invalid token" (type *goa.ServiceError): http.StatusUnauthorized
//   - error: internal error
func DecodeCreateBaseResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body CreateBaseOKResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Base", "CreateBase", err)
			}
			err = ValidateCreateBaseOKResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Base", "CreateBase", err)
			}
			res := NewCreateBaseResultOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body CreateBaseInvalidTokenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Base", "CreateBase", err)
			}
			err = ValidateCreateBaseInvalidTokenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Base", "CreateBase", err)
			}
			return nil, NewCreateBaseInvalidToken(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Base", "CreateBase", resp.StatusCode, string(body))
		}
	}
}

// BuildReceiveDeleteBaseRequest instantiates a HTTP request object with method
// and path set to call the "Base" service "receiveDeleteBase" endpoint
func (c *Client) BuildReceiveDeleteBaseRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		channel string
	)
	{
		p, ok := v.(*base.ReceiveDeleteBasePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("Base", "receiveDeleteBase", "*base.ReceiveDeleteBasePayload", v)
		}
		channel = p.Channel
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ReceiveDeleteBaseBasePath(channel)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Base", "receiveDeleteBase", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeReceiveDeleteBaseRequest returns an encoder for requests sent to the
// Base receiveDeleteBase server.
func EncodeReceiveDeleteBaseRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*base.ReceiveDeleteBasePayload)
		if !ok {
			return goahttp.ErrInvalidType("Base", "receiveDeleteBase", "*base.ReceiveDeleteBasePayload", v)
		}
		if p.JWT != nil {
			head := *p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeReceiveDeleteBaseResponse returns a decoder for responses returned by
// the Base receiveDeleteBase endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeReceiveDeleteBaseResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusSwitchingProtocols:
			res := NewReceiveDeleteBaseDeleteBaseResultSwitchingProtocols()
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Base", "receiveDeleteBase", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteBaseRequest instantiates a HTTP request object with method and
// path set to call the "Base" service "DeleteBase" endpoint
func (c *Client) BuildDeleteBaseRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*base.DeleteBasePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("Base", "DeleteBase", "*base.DeleteBasePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteBaseBasePath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Base", "DeleteBase", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteBaseRequest returns an encoder for requests sent to the Base
// DeleteBase server.
func EncodeDeleteBaseRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*base.DeleteBasePayload)
		if !ok {
			return goahttp.ErrInvalidType("Base", "DeleteBase", "*base.DeleteBasePayload", v)
		}
		if p.JWT != nil {
			head := *p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeDeleteBaseResponse returns a decoder for responses returned by the
// Base DeleteBase endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeDeleteBaseResponse may return the following errors:
//   - "invalid token" (type *goa.ServiceError): http.StatusUnauthorized
//   - "not found" (type *goa.ServiceError): http.StatusNotFound
//   - "permission denied" (type *goa.ServiceError): http.StatusForbidden
//   - error: internal error
func DecodeDeleteBaseResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body DeleteBaseOKResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Base", "DeleteBase", err)
			}
			err = ValidateDeleteBaseOKResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Base", "DeleteBase", err)
			}
			res := NewDeleteBaseResultOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body DeleteBaseInvalidTokenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Base", "DeleteBase", err)
			}
			err = ValidateDeleteBaseInvalidTokenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Base", "DeleteBase", err)
			}
			return nil, NewDeleteBaseInvalidToken(&body)
		case http.StatusNotFound:
			var (
				body DeleteBaseNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Base", "DeleteBase", err)
			}
			err = ValidateDeleteBaseNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Base", "DeleteBase", err)
			}
			return nil, NewDeleteBaseNotFound(&body)
		case http.StatusForbidden:
			var (
				body DeleteBasePermissionDeniedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Base", "DeleteBase", err)
			}
			err = ValidateDeleteBasePermissionDeniedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Base", "DeleteBase", err)
			}
			return nil, NewDeleteBasePermissionDenied(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Base", "DeleteBase", resp.StatusCode, string(body))
		}
	}
}

// BuildReceiveUpdateBaseInfoRequest instantiates a HTTP request object with
// method and path set to call the "Base" service "receiveUpdateBaseInfo"
// endpoint
func (c *Client) BuildReceiveUpdateBaseInfoRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		channel string
	)
	{
		p, ok := v.(*base.ReceiveUpdateBaseInfoPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("Base", "receiveUpdateBaseInfo", "*base.ReceiveUpdateBaseInfoPayload", v)
		}
		channel = p.Channel
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ReceiveUpdateBaseInfoBasePath(channel)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Base", "receiveUpdateBaseInfo", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeReceiveUpdateBaseInfoRequest returns an encoder for requests sent to
// the Base receiveUpdateBaseInfo server.
func EncodeReceiveUpdateBaseInfoRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*base.ReceiveUpdateBaseInfoPayload)
		if !ok {
			return goahttp.ErrInvalidType("Base", "receiveUpdateBaseInfo", "*base.ReceiveUpdateBaseInfoPayload", v)
		}
		if p.JWT != nil {
			head := *p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeReceiveUpdateBaseInfoResponse returns a decoder for responses returned
// by the Base receiveUpdateBaseInfo endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeReceiveUpdateBaseInfoResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusSwitchingProtocols:
			res := NewReceiveUpdateBaseInfoUpdateBaseInfoResultSwitchingProtocols()
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Base", "receiveUpdateBaseInfo", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateBaseInfoRequest instantiates a HTTP request object with method
// and path set to call the "Base" service "UpdateBaseInfo" endpoint
func (c *Client) BuildUpdateBaseInfoRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*base.UpdateBaseInfoPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("Base", "UpdateBaseInfo", "*base.UpdateBaseInfoPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateBaseInfoBasePath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Base", "UpdateBaseInfo", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateBaseInfoRequest returns an encoder for requests sent to the Base
// UpdateBaseInfo server.
func EncodeUpdateBaseInfoRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*base.UpdateBaseInfoPayload)
		if !ok {
			return goahttp.ErrInvalidType("Base", "UpdateBaseInfo", "*base.UpdateBaseInfoPayload", v)
		}
		if p.JWT != nil {
			head := *p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewUpdateBaseInfoRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("Base", "UpdateBaseInfo", err)
		}
		return nil
	}
}

// DecodeUpdateBaseInfoResponse returns a decoder for responses returned by the
// Base UpdateBaseInfo endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeUpdateBaseInfoResponse may return the following errors:
//   - "invalid token" (type *goa.ServiceError): http.StatusUnauthorized
//   - "not found" (type *goa.ServiceError): http.StatusNotFound
//   - "permission denied" (type *goa.ServiceError): http.StatusForbidden
//   - error: internal error
func DecodeUpdateBaseInfoResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateBaseInfoOKResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Base", "UpdateBaseInfo", err)
			}
			err = ValidateUpdateBaseInfoOKResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Base", "UpdateBaseInfo", err)
			}
			res := NewUpdateBaseInfoResultOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body UpdateBaseInfoInvalidTokenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Base", "UpdateBaseInfo", err)
			}
			err = ValidateUpdateBaseInfoInvalidTokenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Base", "UpdateBaseInfo", err)
			}
			return nil, NewUpdateBaseInfoInvalidToken(&body)
		case http.StatusNotFound:
			var (
				body UpdateBaseInfoNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Base", "UpdateBaseInfo", err)
			}
			err = ValidateUpdateBaseInfoNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Base", "UpdateBaseInfo", err)
			}
			return nil, NewUpdateBaseInfoNotFound(&body)
		case http.StatusForbidden:
			var (
				body UpdateBaseInfoPermissionDeniedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Base", "UpdateBaseInfo", err)
			}
			err = ValidateUpdateBaseInfoPermissionDeniedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Base", "UpdateBaseInfo", err)
			}
			return nil, NewUpdateBaseInfoPermissionDenied(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Base", "UpdateBaseInfo", resp.StatusCode, string(body))
		}
	}
}

// BuildReceiveReorderBaseNavStatesRequest instantiates a HTTP request object
// with method and path set to call the "Base" service
// "receiveReorderBaseNavStates" endpoint
func (c *Client) BuildReceiveReorderBaseNavStatesRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		channel string
	)
	{
		p, ok := v.(*base.ReceiveReorderBaseNavStatesPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("Base", "receiveReorderBaseNavStates", "*base.ReceiveReorderBaseNavStatesPayload", v)
		}
		channel = p.Channel
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ReceiveReorderBaseNavStatesBasePath(channel)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Base", "receiveReorderBaseNavStates", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeReceiveReorderBaseNavStatesRequest returns an encoder for requests
// sent to the Base receiveReorderBaseNavStates server.
func EncodeReceiveReorderBaseNavStatesRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*base.ReceiveReorderBaseNavStatesPayload)
		if !ok {
			return goahttp.ErrInvalidType("Base", "receiveReorderBaseNavStates", "*base.ReceiveReorderBaseNavStatesPayload", v)
		}
		if p.JWT != nil {
			head := *p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeReceiveReorderBaseNavStatesResponse returns a decoder for responses
// returned by the Base receiveReorderBaseNavStates endpoint. restoreBody
// controls whether the response body should be restored after having been read.
func DecodeReceiveReorderBaseNavStatesResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusSwitchingProtocols:
			res := NewReceiveReorderBaseNavStatesReorderBaseNavStateResultSwitchingProtocols()
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Base", "receiveReorderBaseNavStates", resp.StatusCode, string(body))
		}
	}
}

// BuildReorderBaseNavStatesRequest instantiates a HTTP request object with
// method and path set to call the "Base" service "ReorderBaseNavStates"
// endpoint
func (c *Client) BuildReorderBaseNavStatesRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ReorderBaseNavStatesBasePath()}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Base", "ReorderBaseNavStates", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeReorderBaseNavStatesRequest returns an encoder for requests sent to
// the Base ReorderBaseNavStates server.
func EncodeReorderBaseNavStatesRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*base.ReorderBaseNavStatesPayload)
		if !ok {
			return goahttp.ErrInvalidType("Base", "ReorderBaseNavStates", "*base.ReorderBaseNavStatesPayload", v)
		}
		if p.JWT != nil {
			head := *p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewReorderBaseNavStatesRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("Base", "ReorderBaseNavStates", err)
		}
		return nil
	}
}

// DecodeReorderBaseNavStatesResponse returns a decoder for responses returned
// by the Base ReorderBaseNavStates endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeReorderBaseNavStatesResponse may return the following errors:
//   - "invalid token" (type *goa.ServiceError): http.StatusUnauthorized
//   - error: internal error
func DecodeReorderBaseNavStatesResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ReorderBaseNavStatesOKResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Base", "ReorderBaseNavStates", err)
			}
			err = ValidateReorderBaseNavStatesOKResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Base", "ReorderBaseNavStates", err)
			}
			res := NewReorderBaseNavStatesReorderBaseNavStateResultOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body ReorderBaseNavStatesInvalidTokenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Base", "ReorderBaseNavStates", err)
			}
			err = ValidateReorderBaseNavStatesInvalidTokenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Base", "ReorderBaseNavStates", err)
			}
			return nil, NewReorderBaseNavStatesInvalidToken(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Base", "ReorderBaseNavStates", resp.StatusCode, string(body))
		}
	}
}

// unmarshalCommandResponseBodyToBaseCommand builds a value of type
// *base.Command from a value of type *CommandResponseBody.
func unmarshalCommandResponseBodyToBaseCommand(v *CommandResponseBody) *base.Command {
	res := &base.Command{
		Type: *v.Type,
	}

	return res
}

// unmarshalCreateBaseResultDataResponseBodyToBaseCreateBaseResultData builds a
// value of type *base.CreateBaseResultData from a value of type
// *CreateBaseResultDataResponseBody.
func unmarshalCreateBaseResultDataResponseBodyToBaseCreateBaseResultData(v *CreateBaseResultDataResponseBody) *base.CreateBaseResultData {
	res := &base.CreateBaseResultData{
		ID: *v.ID,
	}
	res.Info = unmarshalBaseInfoResponseBodyToBaseBaseInfo(v.Info)
	res.NavState = unmarshalBaseNavStateResponseBodyToBaseBaseNavState(v.NavState)

	return res
}

// unmarshalBaseInfoResponseBodyToBaseBaseInfo builds a value of type
// *base.BaseInfo from a value of type *BaseInfoResponseBody.
func unmarshalBaseInfoResponseBodyToBaseBaseInfo(v *BaseInfoResponseBody) *base.BaseInfo {
	res := &base.BaseInfo{
		ID:        *v.ID,
		BaseID:    *v.BaseID,
		Color:     *v.Color,
		Logo:      *v.Logo,
		Name:      *v.Name,
		DeletedAt: v.DeletedAt,
	}

	return res
}

// unmarshalBaseNavStateResponseBodyToBaseBaseNavState builds a value of type
// *base.BaseNavState from a value of type *BaseNavStateResponseBody.
func unmarshalBaseNavStateResponseBodyToBaseBaseNavState(v *BaseNavStateResponseBody) *base.BaseNavState {
	res := &base.BaseNavState{
		ID:        *v.ID,
		BaseID:    *v.BaseID,
		Index:     *v.Index,
		DeletedAt: v.DeletedAt,
	}

	return res
}

// unmarshalDeleteBaseResultDataResponseBodyToBaseDeleteBaseResultData builds a
// value of type *base.DeleteBaseResultData from a value of type
// *DeleteBaseResultDataResponseBody.
func unmarshalDeleteBaseResultDataResponseBodyToBaseDeleteBaseResultData(v *DeleteBaseResultDataResponseBody) *base.DeleteBaseResultData {
	res := &base.DeleteBaseResultData{
		BaseID: *v.BaseID,
	}

	return res
}
