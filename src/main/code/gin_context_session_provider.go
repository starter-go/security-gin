package code

import (
	"context"

	"github.com/starter-go/application/properties"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security"
	"github.com/starter-go/security/jwt"
)

// GinContextSessionProvider ...
type GinContextSessionProvider struct {

	//starter:component
	_as func(security.SessionProvider, security.SessionRegistry) //starter:as(".",".")

	JWTSer jwt.Service //starter:inject("#")

}

func (inst *GinContextSessionProvider) _impl() (security.SessionProvider, security.SessionRegistry) {
	return inst, inst
}

// Registration ...
func (inst *GinContextSessionProvider) Registration() *security.SessionRegistration {
	return &security.SessionRegistration{
		Name:     "GinContextSessionProvider",
		Enabled:  true,
		Priority: 0,
		Provider: inst,
	}
}

// Support ...
func (inst *GinContextSessionProvider) Support(c context.Context) bool {
	h, err := libgin.GetContextHolder(c)
	return (err == nil) && (h != nil)
}

// Current ...
func (inst *GinContextSessionProvider) Current(c context.Context) (security.Session, error) {
	h, err := libgin.GetContextHolder(c)
	if err != nil {
		return nil, err
	}
	session := &ginContextSession{
		context: h.Context(),
		jwts:    inst.JWTSer,
	}
	return session, nil
}

////////////////////////////////////////////////////////////////////////////////

type ginContextSession struct {
	context context.Context
	jwts    jwt.Service
	session *rbac.SessionDTO
}

func (inst *ginContextSession) _impl() security.Session {
	return inst
}

func (inst *ginContextSession) load() *rbac.SessionDTO {
	session := &rbac.SessionDTO{}
	token, err := inst.jwts.GetDTO(inst.context)
	if err == nil {
		*session = token.Session
	}
	return session
}

func (inst *ginContextSession) Get() *rbac.SessionDTO {
	se := inst.session
	if se == nil {
		se = inst.load()
		inst.session = se
	}
	return se
}

func (inst *ginContextSession) Set(s *rbac.SessionDTO) {
	if s == nil {
		s = &rbac.SessionDTO{}
	}
	token := &jwt.Token{}
	token.Session = *s
	err := inst.jwts.SetDTO(inst.context, token)
	if err != nil {
		panic(err)
	}
	inst.session = s
}

func (inst *ginContextSession) UserID() rbac.UserID {
	ses := inst.Get()
	return ses.User.ID
}

func (inst *ginContextSession) UserName() rbac.UserName {
	ses := inst.Get()
	return ses.User.Name
}

func (inst *ginContextSession) Nickname() string {
	ses := inst.Get()
	return ses.User.NickName
}

func (inst *ginContextSession) Avatar() string {
	ses := inst.Get()
	return ses.User.Avatar
}

func (inst *ginContextSession) Roles() rbac.RoleNameList {
	ses := inst.Get()
	return ses.User.Roles
}

func (inst *ginContextSession) Authenticated() bool {
	ses := inst.Get()
	return ses.Authenticated
}

func (inst *ginContextSession) GetProperties() properties.Table {
	ses := inst.Get()
	src := ses.Properties
	dst := properties.NewTable(nil)
	dst.Import(src)
	return dst
}
