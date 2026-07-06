package main4securitygin
import (
    p896fdb913 "github.com/starter-go/security-gin/lib/web/ctrl"
    pfd2c28477 "github.com/starter-go/v0/subjects"
     "github.com/starter-go/application"
)

// type p896fdb913.ContextBindingController in package:github.com/starter-go/security-gin/lib/web/ctrl
//
// id:com-896fdb91368de65f-ctrl-ContextBindingController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p896fdb9136_ctrl_ContextBindingController struct {
}

func (inst* p896fdb9136_ctrl_ContextBindingController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-896fdb91368de65f-ctrl-ContextBindingController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p896fdb9136_ctrl_ContextBindingController) new() any {
    return &p896fdb913.ContextBindingController{}
}

func (inst* p896fdb9136_ctrl_ContextBindingController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p896fdb913.ContextBindingController)
	nop(ie, com)

	
    com.ChainHolder = inst.getChainHolder(ie)
    com.Bypass = inst.getBypass(ie)


    return nil
}


func (inst*p896fdb9136_ctrl_ContextBindingController) getChainHolder(ie application.InjectionExt)pfd2c28477.FilterChainHolder{
    return ie.GetComponent("#alias-fd2c28477d8555ea1fa4190037afa453-FilterChainHolder").(pfd2c28477.FilterChainHolder)
}


func (inst*p896fdb9136_ctrl_ContextBindingController) getBypass(ie application.InjectionExt)bool{
    return ie.GetBool("${security.web.bypass}")
}



// type p896fdb913.GinContextJWTAdapter in package:github.com/starter-go/security-gin/lib/web/ctrl
//
// id:com-896fdb91368de65f-ctrl-GinContextJWTAdapter
// class:class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle
// alias:
// scope:singleton
//
type p896fdb9136_ctrl_GinContextJWTAdapter struct {
}

func (inst* p896fdb9136_ctrl_GinContextJWTAdapter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-896fdb91368de65f-ctrl-GinContextJWTAdapter"
	r.Classes = "class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p896fdb9136_ctrl_GinContextJWTAdapter) new() any {
    return &p896fdb913.GinContextJWTAdapter{}
}

func (inst* p896fdb9136_ctrl_GinContextJWTAdapter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p896fdb913.GinContextJWTAdapter)
	nop(ie, com)

	


    return nil
}



// type p896fdb913.WebRbacController in package:github.com/starter-go/security-gin/lib/web/ctrl
//
// id:com-896fdb91368de65f-ctrl-WebRbacController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p896fdb9136_ctrl_WebRbacController struct {
}

func (inst* p896fdb9136_ctrl_WebRbacController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-896fdb91368de65f-ctrl-WebRbacController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p896fdb9136_ctrl_WebRbacController) new() any {
    return &p896fdb913.WebRbacController{}
}

func (inst* p896fdb9136_ctrl_WebRbacController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p896fdb913.WebRbacController)
	nop(ie, com)

	
    com.Bypass = inst.getBypass(ie)


    return nil
}


func (inst*p896fdb9136_ctrl_WebRbacController) getBypass(ie application.InjectionExt)bool{
    return ie.GetBool("${security.web.bypass}")
}


