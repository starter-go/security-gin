package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/v0/subjects"
	"github.com/starter-go/vlog"
)

// WebRbacController  这个控制器用于配置预先绑定上下文的中间件
type WebRbacController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Bypass bool //starter:inject("${security.web.bypass}")

}

func (inst *WebRbacController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *WebRbacController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Route: inst.route,
	}
}

func (inst *WebRbacController) route(rp libgin.RouterProxy) error {
	routing := &libgin.Routing{
		Priority:   9909,
		Middleware: true,
		Handlers:   []gin.HandlerFunc{inst.handle},
	}
	rp.Route(routing)
	return nil
}

func (inst *WebRbacController) handle(c *gin.Context) {
	err := inst.doCheckSubject(c)
	if err != nil {
		code := http.StatusForbidden
		sender := new(innerWebErrorSender)
		sender.abort(c, code, err)
	}
}

func (inst *WebRbacController) doCheckSubject(c *gin.Context) error {

	method := c.Request.Method
	path := c.Request.Pattern

	vlog.Debug("[request method:'%s' path:'%s']", method, path)

	sub, err := subjects.GetCurrent(c)
	if err != nil {
		return err
	}

	gett, err := sub.DoGet()
	if err != nil {
		return err
	}

	uid := gett.GetUserID()
	roles := gett.GetRoles()

	vlog.Debug("[subject uid:%d roles:'%s']", uid, roles)

	return nil
}
