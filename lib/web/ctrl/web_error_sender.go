package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
)

type innerWebErrorSender struct {
}

func (inst *innerWebErrorSender) abort(c *gin.Context, code int, err error) {

	view := new(rbac.PermissionVO)
	now := lang.Now()

	view.Status = code
	view.Message = http.StatusText(code)
	view.Time = now.Time()
	view.Timestamp = now

	if err != nil {
		view.Error = err.Error()
	}

	c.AbortWithStatusJSON(code, view)
}
