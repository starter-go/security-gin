package code

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgin"
	"github.com/starter-go/security/jwt"
)

// GinContextJWTAdapter ...
type GinContextJWTAdapter struct {

	//starter:component
	_as func(jwt.Registry) //starter:as(".")

	JWTService jwt.Service //starter:inject("#")

	UseCookie  bool  //starter:inject("${security.jwt.use-cookie}")
	UseHeader  bool  //starter:inject("${security.jwt.use-header}")
	MaxAgeInMS int64 //starter:inject("${security.jwt.max-age-in-ms}")
}

func (inst *GinContextJWTAdapter) _impl() (jwt.Adapter, jwt.Registry) {
	return inst, inst
}

// ListRegistrations ...
func (inst *GinContextJWTAdapter) ListRegistrations() []*jwt.Registration {
	r1 := &jwt.Registration{
		Enabled:  true,
		Priority: 0,
		Adapter:  inst,
	}
	return []*jwt.Registration{r1}
}

func (inst *GinContextJWTAdapter) getGinContext(c context.Context) (*gin.Context, error) {
	h, err := libgin.GetContextHolder(c)
	if err != nil {
		return nil, err
	}
	gc := h.GinContext()
	return gc, nil
}

// Accept ...
func (inst *GinContextJWTAdapter) Accept(c context.Context) bool {
	gc, err := inst.getGinContext(c)
	return (err == nil) && (gc != nil)
}

// SetDTO ...
func (inst *GinContextJWTAdapter) SetDTO(c context.Context, o *jwt.Token) error {

	now := lang.Now()
	o.CreatedAt = now
	o.UpdatedAt = now
	o.ExpiredAt = now + lang.Time(inst.MaxAgeInMS)

	text, err := inst.JWTService.Encode(o)
	if err != nil {
		return err
	}
	return inst.SetText(c, text)
}

// SetText ...
func (inst *GinContextJWTAdapter) SetText(c context.Context, t jwt.Text) error {

	value := t.String()
	gc, err := inst.getGinContext(c)
	if err != nil {
		return err
	}

	if inst.UseHeader {
		key1 := inst.key(true)
		gc.Header(key1, value)
	}

	if inst.UseCookie {
		maxAge := int(inst.MaxAgeInMS / 1000)
		path := ""
		domain := ""
		secure := false
		httpOnly := false
		key2 := inst.key(false)
		gc.SetCookie(key2, value, maxAge, path, domain, secure, httpOnly)
	}

	return nil
}

// GetDTO ...
func (inst *GinContextJWTAdapter) GetDTO(c context.Context) (*jwt.Token, error) {
	text, err := inst.GetText(c)
	if err != nil {
		return &jwt.Token{}, nil
	}
	token, err := inst.JWTService.Decode(text)
	if err != nil {
		return &jwt.Token{}, nil
	}
	return token, nil
}

// GetText ...
func (inst *GinContextJWTAdapter) GetText(c context.Context) (jwt.Text, error) {
	gc, err := inst.getGinContext(c)
	if err != nil {
		return "", err
	}
	key := inst.key(false)
	value := ""
	if inst.UseHeader {
		value = gc.GetHeader(key)
	}
	if inst.UseCookie && (value == "") {
		val, _ := gc.Cookie(key)
		value = val
	}
	return jwt.Text(value), nil
}

func (inst *GinContextJWTAdapter) key(setter bool) string {
	if setter {
		return "x-set-jwt"
	}
	return "x-jwt"
}
