package repositories

import "github.com/samber/do/v2"

var Package = do.Package(
	do.Lazy(NewDatabase),
	do.Lazy(NewUserRepository),
)
