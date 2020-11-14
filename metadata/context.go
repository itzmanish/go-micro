// Package metadata is a way of defining message headers
package metadata

import (
	"context"
	"strings"
)

type mdIncomingKey struct{}
type mdOutgoingKey struct{}
type metadataKey struct{}

type rawMetadata struct {
	md Metadata
}

// FromIncomingContext returns metadata from incoming ctx
// returned metadata shoud not be modified or race condition happens
func FromIncomingContext(ctx context.Context) (Metadata, bool) {
	if ctx == nil {
		return nil, false
	}
	md, ok := ctx.Value(mdIncomingKey{}).(*rawMetadata)
	if !ok || md.md == nil {
		return nil, false
	}
	return md.md, ok
}

// FromOutgoingContext returns metadata from outgoing ctx
// returned metadata shoud not be modified or race condition happens
func FromOutgoingContext(ctx context.Context) (Metadata, bool) {
	if ctx == nil {
		return nil, false
	}
	md, ok := ctx.Value(mdOutgoingKey{}).(*rawMetadata)
	if !ok || md.md == nil {
		return nil, false
	}
	return md.md, ok
}

// FromContext returns metadata from the given context
// returned metadata shoud not be modified or race condition happens
//
// Deprecated: use FromIncomingContext or FromOutgoingContext
func FromContext(ctx context.Context) (Metadata, bool) {
	if ctx == nil {
		return nil, false
	}
	md, ok := ctx.Value(metadataKey{}).(Metadata)
	if !ok || md == nil {
		return nil, false
	}
	// capitalise all values
	newMD := make(Metadata, len(md))
	for k, v := range md {
		newMD[strings.Title(k)] = v
	}
	return newMD, ok
}

// NewContext creates a new context with the given metadata
//
// Deprecated: use NewIncomingContext or NewOutgoingContext
func NewContext(ctx context.Context, md Metadata) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = context.WithValue(ctx, metadataKey{}, md)
	ctx = context.WithValue(ctx, mdIncomingKey{}, &rawMetadata{})
	ctx = context.WithValue(ctx, mdOutgoingKey{}, &rawMetadata{})
	return ctx
}

// SetOutgoingContext modify outgoing context with given metadata
func SetOutgoingContext(ctx context.Context, md Metadata) bool {
	if ctx == nil {
		return false
	}
	if omd, ok := ctx.Value(mdOutgoingKey{}).(*rawMetadata); ok {
		omd.md = md
		return true
	}
	return false
}

// SetIncomingContext modify incoming context with given metadata
func SetIncomingContext(ctx context.Context, md Metadata) bool {
	if ctx == nil {
		return false
	}
	if omd, ok := ctx.Value(mdIncomingKey{}).(*rawMetadata); ok {
		omd.md = md
		return true
	}
	return false
}

// NewIncomingContext creates a new context with incoming metadata attached
func NewIncomingContext(ctx context.Context, md Metadata) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = context.WithValue(ctx, mdIncomingKey{}, &rawMetadata{md})
	ctx = context.WithValue(ctx, mdOutgoingKey{}, &rawMetadata{})
	return ctx
}

// NewOutgoingContext creates a new context with outcoming metadata attached
func NewOutgoingContext(ctx context.Context, md Metadata) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = context.WithValue(ctx, mdOutgoingKey{}, &rawMetadata{md})
	ctx = context.WithValue(ctx, mdIncomingKey{}, &rawMetadata{})
	return ctx
}
