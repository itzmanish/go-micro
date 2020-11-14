// Package resolver provides a micro rpc resolver which prefixes a namespace
package api

import (
	"net/http"

	"github.com/itzmanish/go-micro/v2/api/resolver"
	"github.com/itzmanish/go-micro/v2/registry"
)

// default resolver for legacy purposes
// it uses proxy routing to resolve names
// /foo becomes namespace.foo
// /v1/foo becomes namespace.v1.foo
type Resolver struct {
	Options resolver.Options
}

func (r *Resolver) Resolve(req *http.Request, opts ...resolver.ResolveOption) (*resolver.Endpoint, error) {
	// parse options
	options := resolver.NewResolveOptions(opts...)
	var name, method string

	switch r.Options.Handler {
	// internal handlers
	case "meta", "api", "rpc", "micro":
		name, method = apiRoute(req.URL.Path)
	default:
		method = req.Method
		name = proxyRoute(req.URL.Path)
	}
	// append the service prefix, e.g. foo.api
	if len(r.Options.ServicePrefix) > 0 {
		name = r.Options.ServicePrefix + "." + "service." + name
	}

	// check for the namespace in the request header, this can be set by the client or injected
	// by the auth wrapper if an auth token was provided. The headr takes priority over any domain
	// passed as a default
	domain := options.Domain
	if dom := req.Header.Get("Micro-Namespace"); len(dom) > 0 && dom != domain {
		domain = dom
	} else if len(domain) == 0 {
		domain = registry.DefaultDomain
	}

	return &resolver.Endpoint{
		Name:   name,
		Domain: domain,
		Method: method,
	}, nil
}

func (r *Resolver) String() string {
	return "micro"
}

// NewResolver creates a new micro resolver
func NewResolver(opts ...resolver.Option) resolver.Resolver {
	return &Resolver{
		Options: resolver.NewOptions(opts...),
	}
}
