package securitygin

import (
	"github.com/starter-go/application"
	"github.com/starter-go/module-gorm-mysql/modules/mysql"
	securitygin "github.com/starter-go/security-gin"
	"github.com/starter-go/security-gin/gen/main4securitygin"
	"github.com/starter-go/security-gin/gen/test4securitygin"
	"github.com/starter-go/security/modules/security"
	"github.com/starter-go/v0/subjects/modules/subjects"
)

// Module 导出模块 [github.com/starter-go/security-gin]
func Module() application.Module {
	mb := securitygin.NewMainModule()
	mb.Components(main4securitygin.ComForSecurityGorm)
	mb.Depend(security.Module())
	mb.Depend(subjects.ModuleForLib())

	mb.Depend(mysql.Module())
	// mb.Depend(sqlserver.Module())

	return mb.Create()
}

// ModuleForTest ... 导出模块
func ModuleForTest() application.Module {
	mb := securitygin.NewTestModule()
	mb.Components(test4securitygin.ComForSecurityGormTest)
	mb.Depend(Module())
	return mb.Create()
}
