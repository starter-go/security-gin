package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/base/context2"
	"github.com/starter-go/libgin"
	"github.com/starter-go/v0/subjects"
)

// ContextBindingController  这个控制器用于配置预先绑定上下文的中间件
type ContextBindingController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	ChainHolder subjects.FilterChainHolder //starter:inject("#")

	Bypass bool //starter:inject("${security.web.bypass}")

	ctx2adapter context2.Adapter
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
	routing := &libgin.Routing{
		Priority:   9999,
		Middleware: true,
		Handlers:   []gin.HandlerFunc{inst.handle},
	}
	rp.Route(routing)
	return nil
}

func (inst *ContextBindingController) handle(c *gin.Context) {
	err := inst.doSetupContext(c)
	if err != nil {
		code := http.StatusInternalServerError
		sender := new(innerWebErrorSender)
		sender.abort(c, code, err)
	}
}

func (inst *ContextBindingController) innerGetCtx2adapter() context2.Adapter {

	ada := new(innerContextBindingCtx2Adapter)

	return ada
}

func (inst *ContextBindingController) doSetupContext(c1 *gin.Context) error {

	ada1 := inst.innerGetCtx2adapter()
	ctx, err := context2.Setup(c1, func(name string, value *context2.Context) {
		c1.Set(name, value)
		value.Adapter = ada1
	})
	if err != nil {
		return err
	}

	ada2, err := subjects.GetAdapter(ctx)
	if err != nil {
		return err
	}

	holder, err := ada2.GetHolder(ctx)
	if err != nil {
		return err
	}

	sctx := holder.Context
	sctx.ChainHolder = inst.ChainHolder

	sub, err := subjects.GetCurrent(ctx)
	if err != nil {
		return err
	}

	return sub.Load()
}

////////////////////////////////////////////////////////////////////////////////

type innerContextBindingCtx2Adapter struct {
}

// GetValue implements context2.Adapter.
func (inst *innerContextBindingCtx2Adapter) GetValue(c *context2.Context, name any) any {

	raw := c.Raw
	gc := raw.(*gin.Context)
	nameStr := inst.innerGetNameStr(name)

	attr, ok := gc.Get(nameStr)
	if ok {
		return attr
	}
	return nil
}

// SetValue implements context2.Adapter.
func (inst *innerContextBindingCtx2Adapter) SetValue(c *context2.Context, name any, value any) {

	raw := c.Raw
	gc := raw.(*gin.Context)
	nameStr := inst.innerGetNameStr(name)

	gc.Set(nameStr, value)
}

func (i *innerContextBindingCtx2Adapter) innerGetNameStr(name any) string {
	str, ok := name.(string)
	if ok {
		return str
	}
	return "unnamed"
}

////////////////////////////////////////////////////////////////////////////////
