package securitygintest

import (
	"github.com/starter-go/application"

	securitygin1 "github.com/starter-go/security-gin"

	securitygin2 "github.com/starter-go/security-gin/modules/securitygin"

	"github.com/starter-go/security-gin/gen/test4securitygin"
)

// ModuleForTest ... 导出模块
func ModuleForTest() application.Module {
	mb := securitygin1.NewTestModule()
	mb.Components(test4securitygin.ComForSecurityGormTest)
	mb.Depend(securitygin2.Module())
	return mb.Create()
}
