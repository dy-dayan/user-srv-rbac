package handler

import (
	"context"
	srv "github.com/dy-dayan/user-srv-rbac/idl/dayan/user/srv-rbac"
)

type Handler struct {
}

// 添加角色
func (h *Handler) AddRole(ctx context.Context, req *srv.AddRoleReq, rsp *srv.AddRoleResp) error {


	return nil
}

// 添加角色
func (h *Handler) DelRole(ctx context.Context, req *srv.DelRoleReq, rsp *srv.DelRoleResp) error {


	return nil
}

// 添加角色
func (h *Handler) AddAccess(ctx context.Context, req *srv.AddAccessReq, rsp *srv.AddAccessResp) error {


	return nil
}

// 添加角色
func (h *Handler) DelAccess(ctx context.Context, req *srv.DelAccessReq, rsp *srv.DelAccessResp) error {


	return nil
}
// 添加角色
func (h *Handler) AssignRoleAccess(ctx context.Context, req *srv.AssignRoleAccessReq, rsp *srv.AssignRoleAccessResp) error {


	return nil
}
// 添加角色
func (h *Handler) RemoveRoleAccess(ctx context.Context, req *srv.UnassignRoleAccessReq, rsp *srv.UnassignRoleAccessResp) error {


	return nil
}