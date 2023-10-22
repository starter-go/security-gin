package gen4securitygormtest
import (
    pd1a916a20 "github.com/starter-go/libgin"
    p0291972cb "github.com/starter-go/security-gin/src/test/code"
     "github.com/starter-go/application"
)

// type p0291972cb.Demo1controller in package:github.com/starter-go/security-gin/src/test/code
//
// id:com-0291972cb6d389c8-code-Demo1controller
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p0291972cb6_code_Demo1controller struct {
}

func (inst* p0291972cb6_code_Demo1controller) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0291972cb6d389c8-code-Demo1controller"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0291972cb6_code_Demo1controller) new() any {
    return &p0291972cb.Demo1controller{}
}

func (inst* p0291972cb6_code_Demo1controller) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0291972cb.Demo1controller)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)


    return nil
}


func (inst*p0291972cb6_code_Demo1controller) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


