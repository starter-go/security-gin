package code

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgin"
	"github.com/starter-go/security"
	"github.com/starter-go/security/jwt"
	"github.com/starter-go/security/rbac"
)

// ContextBindingController  这个控制器用于配置预先绑定上下文的中间件
type ContextBindingController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	JWTser            jwt.Service             //starter:inject("#")
	SessionService    security.SessionService //starter:inject("#")
	PermissionService rbac.PermissionService  //starter:inject("#")

	GroupNameList string //starter:inject("${security.web.groups}")
}

func (inst *ContextBindingController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *ContextBindingController) Registration() *libgin.ControllerRegistration {
	groups := inst.getGroups()
	return &libgin.ControllerRegistration{
		Groups: groups,
		Route:  inst.route,
	}
}

func (inst *ContextBindingController) getGroups() []string {
	str := inst.GroupNameList
	src := strings.Split(str, ",")
	dst := make([]string, 0)
	for _, item := range src {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		dst = append(dst, item)
	}
	return dst
}

func (inst *ContextBindingController) route(rp libgin.RouterProxy) error {
	rp.Route(&libgin.Routing{
		Priority:   1000,
		Middleware: true,
		Handlers:   []gin.HandlerFunc{inst.doBind},
	})
	return nil
}

func (inst *ContextBindingController) doBind(c *gin.Context) {

	libgin.BindContext(c)

	sub, err := security.SetupSubject(c, inst.SessionService)
	if err != nil {
		panic(err)
	}

	err = inst.checkPermission(c, sub)
	if err != nil {
		inst.sendNoPermission(c)
		return
	}
}

func (inst *ContextBindingController) checkPermission(c *gin.Context, sub security.Subject) error {

	cache := inst.PermissionService.GetCache()
	perm := &rbac.PermissionDTO{}
	perm.Method = c.Request.Method
	perm.Path = c.FullPath()
	perm2, err := cache.Find(c, perm)
	if err != nil {
		return err
	}

	authenticated := sub.GetSession(true).Authenticated()
	want := perm2.AcceptRoles.List()
	for _, role := range want {
		if authenticated {
			if sub.HasRole(role) {
				return nil
			}
		} else {
			if role == rbac.RoleAny || role == rbac.RoleAnonym {
				return nil
			}
		}
	}
	return fmt.Errorf("no permission")
}

func (inst *ContextBindingController) sendNoPermission(c *gin.Context) {
	now := time.Now()
	code := http.StatusForbidden
	method := c.Request.Method
	path := c.FullPath()
	js := &rbac.BaseVO{}
	js.Status = code
	js.Message = http.StatusText(code)
	js.Error = fmt.Sprintf("no permission to access HTTP.%s(%s)", method, path)
	js.Time = now
	js.Timestamp = lang.NewTime(now)
	c.AbortWithStatusJSON(code, js)
}
