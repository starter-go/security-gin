package main

import (
	"os"

	"github.com/starter-go/security-gin/modules/securitygin"
	"github.com/starter-go/starter"
)

func main() {
	m := securitygin.ModuleForTest()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
