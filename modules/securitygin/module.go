package securitygin

import (
	"github.com/starter-go/application"
	securitygin "github.com/starter-go/security-gin"
	gen4securitygin "github.com/starter-go/security-gin/gen/gen4securitygin"
)

// Module ... 导出模块
func Module() application.Module {
	mb := securitygin.ModuleT()
	mb.Components(gen4securitygin.ComForSecurityGorm)
	return mb.Create()
}
