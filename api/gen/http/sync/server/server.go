// Code generated by goa v3.16.2, DO NOT EDIT.
//
// Sync HTTP server
//
// Command:
// $ goa gen mai.today/api/design

package server

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
	sync "mai.today/api/gen/sync"
)

// Server lists the Sync service endpoint HTTP handlers.
type Server struct {
	Mounts      []*MountPoint
	ReceiveSync http.Handler
	Sync        http.Handler
	CORS        http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the Sync service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *sync.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"ReceiveSync", "GET", "/{channel}/sync"},
			{"Sync", "POST", "/sync"},
			{"CORS", "OPTIONS", "/{channel}/sync"},
			{"CORS", "OPTIONS", "/sync"},
		},
		ReceiveSync: NewReceiveSyncHandler(e.ReceiveSync, mux, decoder, encoder, errhandler, formatter),
		Sync:        NewSyncHandler(e.Sync, mux, decoder, encoder, errhandler, formatter),
		CORS:        NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "Sync" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.ReceiveSync = m(s.ReceiveSync)
	s.Sync = m(s.Sync)
	s.CORS = m(s.CORS)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return sync.MethodNames[:] }

// Mount configures the mux to serve the Sync endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountReceiveSyncHandler(mux, h.ReceiveSync)
	MountSyncHandler(mux, h.Sync)
	MountCORSHandler(mux, h.CORS)
}

// Mount configures the mux to serve the Sync endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountReceiveSyncHandler configures the mux to serve the "Sync" service
// "receiveSync" endpoint.
func MountReceiveSyncHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleSyncOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/{channel}/sync", f)
}

// NewReceiveSyncHandler creates a HTTP handler which loads the HTTP request
// and calls the "Sync" service "receiveSync" endpoint.
func NewReceiveSyncHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeReceiveSyncRequest(mux, decoder)
		encodeResponse = EncodeReceiveSyncResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "receiveSync")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Sync")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountSyncHandler configures the mux to serve the "Sync" service "sync"
// endpoint.
func MountSyncHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleSyncOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/sync", f)
}

// NewSyncHandler creates a HTTP handler which loads the HTTP request and calls
// the "Sync" service "sync" endpoint.
func NewSyncHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeSyncRequest(mux, decoder)
		encodeResponse = EncodeSyncResponse(encoder)
		encodeError    = EncodeSyncError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "sync")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Sync")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service Sync.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleSyncOrigin(h)
	mux.Handle("OPTIONS", "/{channel}/sync", h.ServeHTTP)
	mux.Handle("OPTIONS", "/sync", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 204 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	})
}

// HandleSyncOrigin applies the CORS response headers corresponding to the
// origin for the service Sync.
func HandleSyncOrigin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOrigin(origin, "*") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, PATCH, DELETE, CONNECT, OPTIONS, TRACE")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
				w.WriteHeader(204)
				return
			}
			h.ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
		return
	})
}