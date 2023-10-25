package code

import (
	"net/http"
	"time"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgin"
	"github.com/starter-go/libgin/web"
	"github.com/starter-go/security/rbac"
)

// SecurityGinResponder ...
type SecurityGinResponder struct {

	//starter:component
	_as func(libgin.ResponderRegistry, libgin.Responder) //starter:as(".",".")
}

func (inst *SecurityGinResponder) _impl() (libgin.ResponderRegistry, libgin.Responder) {
	return inst, inst
}

// ListRegistrations ...
func (inst *SecurityGinResponder) ListRegistrations() []*libgin.ResponderRegistration {
	r1 := &libgin.ResponderRegistration{
		Enabled:   true,
		Priority:  0,
		Name:      "security-gin-rbac-responder",
		Responder: inst,
	}
	return []*libgin.ResponderRegistration{r1}
}

// Accept ...
func (inst *SecurityGinResponder) Accept(resp *libgin.Response) bool {
	return true
}

// Send ...
func (inst *SecurityGinResponder) Send(resp *libgin.Response) {

	now := time.Now()
	ctx := resp.Context
	err := resp.Error
	data := resp.Data
	status := resp.Status

	werr, ok := err.(web.Error)
	if ok {
		status = werr.Status()
	}

	if status == 0 {
		if err == nil {
			status = http.StatusOK
		} else {
			status = http.StatusInternalServerError
		}
	}

	vg, ok := data.(rbac.VOGetter)
	if ok {
		v := vg.GetVO()
		if err != nil {
			v.Error = err.Error()
		}
		v.Message = http.StatusText(status)
		v.Status = status
		v.Time = now
		v.Timestamp = lang.NewTime(now)
	}

	ctx.JSON(status, data)
}
