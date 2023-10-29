package code

import (
	"context"
	"fmt"

	"github.com/starter-go/security/rbac"
)

// MockPermissionService ...
type MockPermissionService struct {

	//starter:component
	_as func(rbac.PermissionService) //starter:as("#")

	cache myMockPermissionCache
}

func (inst *MockPermissionService) _impl() {
	inst._as(inst)
}

func (inst *MockPermissionService) Insert(c context.Context, o *rbac.PermissionDTO) (*rbac.PermissionDTO, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *MockPermissionService) Update(c context.Context, id rbac.PermissionID, o *rbac.PermissionDTO) (*rbac.PermissionDTO, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *MockPermissionService) Delete(c context.Context, id rbac.PermissionID) error {
	return fmt.Errorf("no impl")
}

func (inst *MockPermissionService) Find(c context.Context, id rbac.PermissionID) (*rbac.PermissionDTO, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *MockPermissionService) List(c context.Context, q *rbac.PermissionQuery) ([]*rbac.PermissionDTO, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *MockPermissionService) ListAll(c context.Context) ([]*rbac.PermissionDTO, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *MockPermissionService) GetCache() rbac.PermissionCache {
	return &inst.cache
}

////////////////////////////////////////////////////////////////////////////////

type myMockPermissionCache struct{}

func (inst *myMockPermissionCache) _impl() rbac.PermissionCache {
	return inst
}

func (inst *myMockPermissionCache) Clear() {}

func (inst *myMockPermissionCache) Find(c context.Context, want *rbac.PermissionDTO) (*rbac.PermissionDTO, error) {
	have := &rbac.PermissionDTO{}
	have.Method = want.Method
	have.Path = want.Path
	have.AcceptRoles = rbac.NewRoleNameList(rbac.RoleAny, rbac.RoleAnonym, rbac.RoleUser)
	return have, nil
}
