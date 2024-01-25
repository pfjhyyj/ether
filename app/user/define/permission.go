package define

import "github.com/pfjhyyj/ether/common"

type CreatePermissionRequest struct {
	Action      string `json:"action" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdatePermissionRequest struct {
	PermissionId uint   `uri:"permissionId" binding:"required"`
	Action       string `json:"action"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}

type DeletePermissionRequest struct {
	PermissionId uint `uri:"permissionId" binding:"required"`
}

type ListPermissionsRequest struct {
	common.PageRequest
}

type PermissionPageResponse struct {
	PermissionId uint   `json:"permissionId"`
	Action       string `json:"action"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}
