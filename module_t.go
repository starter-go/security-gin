package securitygin

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/security/modules/security"
)

const (
	theModuleName        = "github.com/starter-go/security-gin"
	theModuleVersion     = "v0.0.13"
	theModuleRevision    = 13
	theModuleResPath     = "src/main/resources"
	theTestModuleResPath = "src/test/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

// ModuleT ...
func ModuleT() *application.ModuleBuilder {
	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theModuleResFS, theModuleResPath)

	mb.Depend(security.Module())
	mb.Depend(libgin.Module())
	// mb.Depend(libgorm.Module())
	// mb.Depend(modulegormsqlserver.Module())
	// mb.Depend(modulegormmysql.Module())

	return mb
}

// TestModuleT ...
func TestModuleT() *application.ModuleBuilder {
	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)
	return mb
}
