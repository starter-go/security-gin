package ctrl

import "github.com/starter-go/application"

// GinContextJWTAdapter ...
type GinContextJWTAdapter struct {

	//starter:component

	_as func(application.Lifecycle) //starter:as(".")

	// JWTService jwt.Service //starter:inject("#")

	// UseCookie  bool  //starter:inject("${security.jwt.use-cookie}")
	// UseHeader  bool  //starter:inject("${security.jwt.use-header}")
	// MaxAgeInMS int64 //starter:inject("${security.jwt.max-age-in-ms}")

}

// Life implements application.Lifecycle.
func (inst *GinContextJWTAdapter) Life() *application.Life {
	return new(application.Life)
}

func (inst *GinContextJWTAdapter) _impl() application.Lifecycle {
	return inst
}
