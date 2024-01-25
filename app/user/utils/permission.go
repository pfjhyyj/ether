package utils

import (
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/model"
	"github.com/pfjhyyj/ether/common"
)

func ConvertCreatePermissionRequestToPermission(req *define.CreatePermissionRequest) *model.Permission {
	return &model.Permission{
		Name:        req.Name,
		Action:      req.Action,
		Description: req.Description,
	}
}

func ConvertUpdatePermissionRequestToPermission(req *define.UpdatePermissionRequest) *model.Permission {
	return &model.Permission{
		PermissionId: req.PermissionId,
		Name:         req.Name,
		Action:       req.Action,
		Description:  req.Description,
	}
}

func ConvertListPermissionRequestToParam(req *define.ListPermissionsRequest) *model.QueryPermissionParams {
	return &model.QueryPermissionParams{
		PageRequest: common.PageRequest{
			Current:  req.Current,
			PageSize: req.PageSize,
		},
	}
}

func ConvertPermissionListToPageResponse(permissions []*model.Permission) []*define.PermissionPageResponse {
	var permissionInfos []*define.PermissionPageResponse
	for _, permission := range permissions {
		permissionInfos = append(permissionInfos, &define.PermissionPageResponse{
			PermissionId: permission.PermissionId,
			Name:         permission.Name,
			Action:       permission.Action,
			Description:  permission.Description,
		})
	}
	return permissionInfos
}
