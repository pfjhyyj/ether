use sea_orm::{ColumnTrait, EntityTrait, PaginatorTrait, QueryFilter, QueryOrder, QuerySelect};
use utils::{request::parse_page_and_size, response::{ApiError, PageResponse}};
use domain::entity::{permission, role_permission};

use crate::modules::role::controller::list_permissions::{ListRolePermissionsRequest, ListRolePermissionsResponse};

pub async fn get_page_role_permissions(
    role_id: i64,
    req: ListRolePermissionsRequest
) -> Result<PageResponse<ListRolePermissionsResponse>, ApiError> {
    let db = utils::db::conn();
    
    let query = role_permission::Entity::find()
        .filter(role_permission::Column::RoleId.eq(role_id));
        
    let (offset, limit) = parse_page_and_size(req.page, req.size);

    let total = query.clone().count(db).await.map_err(|e| {
        tracing::error!(error = ?e, "Failed to count role permission");
        ApiError::DbError(None)
    })?;

    let permissions: Vec<(role_permission::Model, Option<permission::Model>)> = query
        .order_by_asc(permission::Column::PermissionId)
        .limit(limit)
        .offset(offset)
        .find_also_related(permission::Entity)
        .all(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to query role permission");
            ApiError::DbError(None)
        })?;

    let resp = PageResponse {
        total,
        page: offset / limit + 1,
        size: limit,
        data: permissions.into_iter().filter_map(|(role_permission, permission)| {
            permission.as_ref().map(|perm| ListRolePermissionsResponse {
                permission_id: role_permission.permission_id,
                object: perm.object.clone(),
                action: perm.action.clone(),
                name: perm.name.clone(),
                description: perm.description.clone(),
            })
        }).collect(),
    };

    Ok(resp)
}