package main4securitygin
import (
    p24287f458 "github.com/starter-go/rbac"
    pd4e0ee677 "github.com/starter-go/security"
    p6d96d35d0 "github.com/starter-go/security-gin/src/main/code"
    p91f218d46 "github.com/starter-go/security/jwt"
    paff1180b7 "github.com/starter-go/security/subjects"
     "github.com/starter-go/application"
)

// type p6d96d35d0.ContextBindingController in package:github.com/starter-go/security-gin/src/main/code
//
// id:com-6d96d35d0126875b-code-ContextBindingController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p6d96d35d01_code_ContextBindingController struct {
}

func (inst* p6d96d35d01_code_ContextBindingController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6d96d35d0126875b-code-ContextBindingController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6d96d35d01_code_ContextBindingController) new() any {
    return &p6d96d35d0.ContextBindingController{}
}

func (inst* p6d96d35d01_code_ContextBindingController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6d96d35d0.ContextBindingController)
	nop(ie, com)

	
    com.JWTser = inst.getJWTser(ie)
    com.SessionService = inst.getSessionService(ie)
    com.PermissionService = inst.getPermissionService(ie)
    com.SubjectsLoader = inst.getSubjectsLoader(ie)
    com.GroupNameList = inst.getGroupNameList(ie)
    com.Bypass = inst.getBypass(ie)


    return nil
}


func (inst*p6d96d35d01_code_ContextBindingController) getJWTser(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}


func (inst*p6d96d35d01_code_ContextBindingController) getSessionService(ie application.InjectionExt)pd4e0ee677.SessionService{
    return ie.GetComponent("#alias-d4e0ee677c339b7ffcf1d55767953499-SessionService").(pd4e0ee677.SessionService)
}


func (inst*p6d96d35d01_code_ContextBindingController) getPermissionService(ie application.InjectionExt)p24287f458.PermissionService{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-PermissionService").(p24287f458.PermissionService)
}


func (inst*p6d96d35d01_code_ContextBindingController) getSubjectsLoader(ie application.InjectionExt)paff1180b7.Loader{
    return ie.GetComponent("#alias-aff1180b734cd089659a2dcc3be458d7-Loader").(paff1180b7.Loader)
}


func (inst*p6d96d35d01_code_ContextBindingController) getGroupNameList(ie application.InjectionExt)string{
    return ie.GetString("${security.web.groups}")
}


func (inst*p6d96d35d01_code_ContextBindingController) getBypass(ie application.InjectionExt)bool{
    return ie.GetBool("${security.web.bypass}")
}



// type p6d96d35d0.GinContextJWTAdapter in package:github.com/starter-go/security-gin/src/main/code
//
// id:com-6d96d35d0126875b-code-GinContextJWTAdapter
// class:class-91f218d46ec21cd234778bbe54aecc66-Registry
// alias:
// scope:singleton
//
type p6d96d35d01_code_GinContextJWTAdapter struct {
}

func (inst* p6d96d35d01_code_GinContextJWTAdapter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6d96d35d0126875b-code-GinContextJWTAdapter"
	r.Classes = "class-91f218d46ec21cd234778bbe54aecc66-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6d96d35d01_code_GinContextJWTAdapter) new() any {
    return &p6d96d35d0.GinContextJWTAdapter{}
}

func (inst* p6d96d35d01_code_GinContextJWTAdapter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6d96d35d0.GinContextJWTAdapter)
	nop(ie, com)

	
    com.JWTService = inst.getJWTService(ie)
    com.UseCookie = inst.getUseCookie(ie)
    com.UseHeader = inst.getUseHeader(ie)
    com.MaxAgeInMS = inst.getMaxAgeInMS(ie)


    return nil
}


func (inst*p6d96d35d01_code_GinContextJWTAdapter) getJWTService(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}


func (inst*p6d96d35d01_code_GinContextJWTAdapter) getUseCookie(ie application.InjectionExt)bool{
    return ie.GetBool("${security.jwt.use-cookie}")
}


func (inst*p6d96d35d01_code_GinContextJWTAdapter) getUseHeader(ie application.InjectionExt)bool{
    return ie.GetBool("${security.jwt.use-header}")
}


func (inst*p6d96d35d01_code_GinContextJWTAdapter) getMaxAgeInMS(ie application.InjectionExt)int64{
    return ie.GetInt64("${security.jwt.max-age-in-ms}")
}


