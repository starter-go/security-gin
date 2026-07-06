package test4securitygin
import (
    pd1a916a20 "github.com/starter-go/libgin"
    p512a30914 "github.com/starter-go/libgorm"
    pbb42fed2b "github.com/starter-go/security-gin/src/test/golang/testcom"
     "github.com/starter-go/application"
)

// type pbb42fed2b.Demo1controller in package:github.com/starter-go/security-gin/src/test/golang/testcom
//
// id:com-bb42fed2be2c7d19-testcom-Demo1controller
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pbb42fed2be_testcom_Demo1controller struct {
}

func (inst* pbb42fed2be_testcom_Demo1controller) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-bb42fed2be2c7d19-testcom-Demo1controller"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pbb42fed2be_testcom_Demo1controller) new() any {
    return &pbb42fed2b.Demo1controller{}
}

func (inst* pbb42fed2be_testcom_Demo1controller) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pbb42fed2b.Demo1controller)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.DataGroups = inst.getDataGroups(ie)


    return nil
}


func (inst*pbb42fed2be_testcom_Demo1controller) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pbb42fed2be_testcom_Demo1controller) getDataGroups(ie application.InjectionExt)[]p512a30914.GroupRegistry{
    dst := make([]p512a30914.GroupRegistry, 0)
    src := ie.ListComponents(".class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry")
    for _, item1 := range src {
        item2 := item1.(p512a30914.GroupRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type pbb42fed2b.MockPermissionService in package:github.com/starter-go/security-gin/src/test/golang/testcom
//
// id:com-bb42fed2be2c7d19-testcom-MockPermissionService
// class:
// alias:alias-24287f4589fe5add27fb48a88d706565-PermissionService
// scope:singleton
//
type pbb42fed2be_testcom_MockPermissionService struct {
}

func (inst* pbb42fed2be_testcom_MockPermissionService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-bb42fed2be2c7d19-testcom-MockPermissionService"
	r.Classes = ""
	r.Aliases = "alias-24287f4589fe5add27fb48a88d706565-PermissionService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pbb42fed2be_testcom_MockPermissionService) new() any {
    return &pbb42fed2b.MockPermissionService{}
}

func (inst* pbb42fed2be_testcom_MockPermissionService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pbb42fed2b.MockPermissionService)
	nop(ie, com)

	


    return nil
}


