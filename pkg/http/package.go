package http

import (
	"github.com/samber/do/v2"
)

// Package provides HTTP-related services for dependency injection
// This demonstrates how to organize related services in a do package.
var Package = do.Package(
	do.Lazy(NewHTTPServer),
	do.Lazy(NewUserHandler),
	do.Lazy(NewHealthHandler),
)
