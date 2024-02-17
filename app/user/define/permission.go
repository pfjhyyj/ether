package define

import "github.com/pfjhyyj/ether/common"

type CreatePermissionRequest struct {
	Target      string `json:"target" binding:"required"`
	Action      string `json:"action" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdatePermissionRequest struct {
	PermissionId uint   `uri:"permissionId" binding:"required"`
	Target       string `json:"target"`
	Action       string `json:"action"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}

type DeletePermissionRequest struct {
	PermissionId uint `uri:"permissionId" binding:"required"`
}

type ListPermissionsRequest struct {
	common.PageRequest
	Name   string `form:"name"`
	Target string `form:"target"`
}

type PermissionPageResponse struct {
	PermissionId uint   `json:"permissionId"`
	Target       string `json:"target"`
	Action       string `json:"action"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}
