// Code generated by goa v3.16.2, DO NOT EDIT.
//
// Base endpoints
//
// Command:
// $ goa gen mai.today/api/design

package base

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "Base" service endpoints.
type Endpoints struct {
	ReceiveCreateBase           goa.Endpoint
	CreateBase                  goa.Endpoint
	ReceiveDeleteBase           goa.Endpoint
	DeleteBase                  goa.Endpoint
	ReceiveUpdateBaseInfo       goa.Endpoint
	UpdateBaseInfo              goa.Endpoint
	ReceiveReorderBaseNavStates goa.Endpoint
	ReorderBaseNavStates        goa.Endpoint
}

// NewEndpoints wraps the methods of the "Base" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		ReceiveCreateBase:           NewReceiveCreateBaseEndpoint(s, a.JWTAuth),
		CreateBase:                  NewCreateBaseEndpoint(s, a.JWTAuth),
		ReceiveDeleteBase:           NewReceiveDeleteBaseEndpoint(s, a.JWTAuth),
		DeleteBase:                  NewDeleteBaseEndpoint(s, a.JWTAuth),
		ReceiveUpdateBaseInfo:       NewReceiveUpdateBaseInfoEndpoint(s, a.JWTAuth),
		UpdateBaseInfo:              NewUpdateBaseInfoEndpoint(s, a.JWTAuth),
		ReceiveReorderBaseNavStates: NewReceiveReorderBaseNavStatesEndpoint(s, a.JWTAuth),
		ReorderBaseNavStates:        NewReorderBaseNavStatesEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "Base" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.ReceiveCreateBase = m(e.ReceiveCreateBase)
	e.CreateBase = m(e.CreateBase)
	e.ReceiveDeleteBase = m(e.ReceiveDeleteBase)
	e.DeleteBase = m(e.DeleteBase)
	e.ReceiveUpdateBaseInfo = m(e.ReceiveUpdateBaseInfo)
	e.UpdateBaseInfo = m(e.UpdateBaseInfo)
	e.ReceiveReorderBaseNavStates = m(e.ReceiveReorderBaseNavStates)
	e.ReorderBaseNavStates = m(e.ReorderBaseNavStates)
}

// NewReceiveCreateBaseEndpoint returns an endpoint function that calls the
// method "receiveCreateBase" of service "Base".
func NewReceiveCreateBaseEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ReceiveCreateBasePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "JWT",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.ReceiveCreateBase(ctx, p)
	}
}

// NewCreateBaseEndpoint returns an endpoint function that calls the method
// "CreateBase" of service "Base".
func NewCreateBaseEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CreateBasePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "JWT",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.CreateBase(ctx, p)
	}
}

// NewReceiveDeleteBaseEndpoint returns an endpoint function that calls the
// method "receiveDeleteBase" of service "Base".
func NewReceiveDeleteBaseEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ReceiveDeleteBasePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "JWT",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.ReceiveDeleteBase(ctx, p)
	}
}

// NewDeleteBaseEndpoint returns an endpoint function that calls the method
// "DeleteBase" of service "Base".
func NewDeleteBaseEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DeleteBasePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "JWT",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.DeleteBase(ctx, p)
	}
}

// NewReceiveUpdateBaseInfoEndpoint returns an endpoint function that calls the
// method "receiveUpdateBaseInfo" of service "Base".
func NewReceiveUpdateBaseInfoEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ReceiveUpdateBaseInfoPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "JWT",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.ReceiveUpdateBaseInfo(ctx, p)
	}
}

// NewUpdateBaseInfoEndpoint returns an endpoint function that calls the method
// "UpdateBaseInfo" of service "Base".
func NewUpdateBaseInfoEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UpdateBaseInfoPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "JWT",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.UpdateBaseInfo(ctx, p)
	}
}

// NewReceiveReorderBaseNavStatesEndpoint returns an endpoint function that
// calls the method "receiveReorderBaseNavStates" of service "Base".
func NewReceiveReorderBaseNavStatesEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ReceiveReorderBaseNavStatesPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "JWT",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.ReceiveReorderBaseNavStates(ctx, p)
	}
}

// NewReorderBaseNavStatesEndpoint returns an endpoint function that calls the
// method "ReorderBaseNavStates" of service "Base".
func NewReorderBaseNavStatesEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ReorderBaseNavStatesPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "JWT",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.ReorderBaseNavStates(ctx, p)
	}
}
