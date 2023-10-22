package main

import (
	"os"

	"github.com/starter-go/security-gin/modules/securitygintest"
	"github.com/starter-go/starter"
)

func main() {
	m := securitygintest.ModuleForTest()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
