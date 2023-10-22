package securitygintest

import (
	"github.com/starter-go/application"

	securitygin1 "github.com/starter-go/security-gin"

	securitygin2 "github.com/starter-go/security-gin/modules/securitygin"

	gen4securitygintest "github.com/starter-go/security-gin/gen/gen4securitygintest"
)

// ModuleForTest ... 导出模块
func ModuleForTest() application.Module {
	mb := securitygin1.TestModuleT()
	mb.Components(gen4securitygintest.ComForSecurityGormTest)
	mb.Depend(securitygin2.Module())
	return mb.Create()
}
