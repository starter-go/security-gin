package code

import (
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/security"
	"github.com/starter-go/security/jwt"
)

// ContextBindingController  这个控制器用于配置预先绑定上下文的中间件
type ContextBindingController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	JWTser         jwt.Service             //starter:inject("#")
	SessionService security.SessionService //starter:inject("#")

}

func (inst *ContextBindingController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *ContextBindingController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Route: inst.route,
	}
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

	_, err := security.SetupSubject(c, inst.SessionService)
	if err != nil {
		panic(err)
	}
}
