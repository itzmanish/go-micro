package resolver

import (
	"net/http"

	"github.com/itzmanish/go-micro/v2/registry"
)

type Options struct {
	Handler       string
	Namespace     func(*http.Request) string
	ServicePrefix string
}

type Option func(o *Options)

// NewOptions returns new initialised options
func NewOptions(opts ...Option) Options {
	var options Options
	for _, o := range opts {
		o(&options)
	}

	if options.Namespace == nil {
		options.Namespace = StaticNamespace("go.micro")
	}

	return options
}

// WithHandler sets the handler being used
func WithHandler(h string) Option {
	return func(o *Options) {
		o.Handler = h
	}
}

// WithNamespace sets the function which determines the namespace for a request
func WithNamespace(n func(*http.Request) string) Option {
	return func(o *Options) {
		o.Namespace = n
	}
}

// WithServicePrefix sets the ServicePrefix option
func WithServicePrefix(p string) Option {
	return func(o *Options) {
		o.ServicePrefix = p
	}
}

// ResolveOptions are used when resolving a request
type ResolveOptions struct {
	Domain string
}

// ResolveOption sets an option
type ResolveOption func(*ResolveOptions)

// Domain sets the resolve Domain option
func Domain(n string) ResolveOption {
	return func(o *ResolveOptions) {
		o.Domain = n
	}
}

// NewResolveOptions returns new initialised resolve options
func NewResolveOptions(opts ...ResolveOption) ResolveOptions {
	var options ResolveOptions
	for _, o := range opts {
		o(&options)
	}
	if len(options.Domain) == 0 {
		options.Domain = registry.DefaultDomain
	}

	return options
}
