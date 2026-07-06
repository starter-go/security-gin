package testcom

import (
	"github.com/gin-gonic/gin"
	"github.com/starter-go/application"
	"github.com/starter-go/libgin"
	"github.com/starter-go/libgorm"
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/subjects"
)

// Demo1controller ...
type Demo1controller struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")

	DataGroups []libgorm.GroupRegistry //starter:inject(".")
}

func (inst *Demo1controller) _impl() libgin.Controller {
	return inst
}

func (inst *Demo1controller) Life() *application.Life {
	return &application.Life{
		// OnStartPre: inst.onStart,
	}
}

// func (inst *Demo1controller) onStart() error {
// 	src := inst.DataGroups
// 	for _, r1 := range src {
// 		tmp := r1.Groups()
// 		for _, r2 := range tmp {
// 			uri := r2.URI
// 			vlog.Debug("data.group.uri = %s", uri)
// 		}
// 	}
// 	return nil
// }

// Registration ...
func (inst *Demo1controller) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Groups: []string{"rest"},
		Route:  inst.route,
	}
}

func (inst *Demo1controller) route(rp libgin.RouterProxy) error {
	rp = rp.For("demo1")
	rp.GET("", inst.handleGet)
	rp.GET(":id", inst.handleGet)
	rp.GET("/demo/:uuid/:name/detail", inst.handleGet)
	return nil
}

func (inst *Demo1controller) handleGet(c *gin.Context) {
	req := &demo1request{
		controller: inst,
		context:    c,
	}
	req.execute(req.do1)
}

////////////////////////////////////////////////////////////////////////////////

// Demo1ID ...
type Demo1ID int

// Demo1vo ...
type Demo1vo struct {
	ID Demo1ID

	rbac.BaseVO

	A   int
	B   string
	Map map[string]string
}

////////////////////////////////////////////////////////////////////////////////

type demo1request struct {
	context    *gin.Context
	controller *Demo1controller

	wantRequestID   bool
	wantRequestBody bool
	wantRequestPage bool

	page         rbac.Pagination
	id           Demo1ID
	requireRoles []rbac.RoleName

	body1 Demo1vo
	body2 Demo1vo
}

func (inst *demo1request) open() error {

	if inst.wantRequestID {
	}

	if inst.wantRequestBody {
	}

	return nil
}

func (inst *demo1request) send(err error) {
	status := inst.body2.Status
	data := &inst.body2
	resp := &libgin.Response{
		Error:   err,
		Context: inst.context,
		Status:  status,
		Data:    data,
	}
	inst.controller.Responder.Send(resp)
}

func (inst *demo1request) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *demo1request) do1() error {

	ctx := inst.context
	sub, err := subjects.GetCurrent(ctx)
	if err != nil {
		return err
	}

	// gett, err := sub.DoGet()

	sett, err := sub.DoSet()
	if err != nil {
		return err
	}

	// session := sub.GetSession()
	// val := fmt.Sprintf("%s", inst.id)
	// session.SetProperty("id", val)
	// session.Create()
	// return session.Commit()

	sett.SetProperty("foo", "11")
	sett.SetProperty("bar", "22")

	sub.Create()

	return sub.Flush()
}
