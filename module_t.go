package securitygin

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/security/modules/security"
)

const (
	theModuleName     = "github.com/starter-go/security-gin"
	theModuleVersion  = "v1.0.50"
	theModuleRevision = 20
)

////////////////////////////////////////////////////////////////////////////////

const (
	theMainModuleResPath = "src/main/resources"
	theTestModuleResPath = "src/test/resources"
)

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

// NewMainModule 导出模块 [github.com/starter-go/security-gin]
func NewMainModule() *application.ModuleBuilder {

	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)

	mb.Depend(security.Module())
	mb.Depend(libgin.Module())

	return mb
}

// NewTestModule 导出模块 [github.com/starter-go/security-gin#test]
func NewTestModule() *application.ModuleBuilder {

	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)

	return mb
}
