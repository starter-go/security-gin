package securitygin

import (
	"github.com/starter-go/application"
	securitygin "github.com/starter-go/security-gin"
	"github.com/starter-go/security-gin/gen/main4securitygin"
	"github.com/starter-go/security/modules/security"
)

// Module 导出模块 [github.com/starter-go/security-gin]
func Module() application.Module {
	mb := securitygin.NewMainModule()
	mb.Components(main4securitygin.ComForSecurityGorm)
	mb.Depend(security.Module())
	return mb.Create()
}
