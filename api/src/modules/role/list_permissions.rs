use salvo::{oapi::extract::{PathParam, QueryParam}, prelude::*};
use sea_orm::{ColumnTrait, EntityTrait, PaginatorTrait, QueryFilter, QueryOrder, QuerySelect};
use serde::{Deserialize, Serialize};
use utils::{request::{parse_page_request, PageRequest}, response::{ApiError, ApiOk, ApiResult, PageResponse}};


#[derive(Debug, Deserialize, ToSchema)]
pub struct ListRolePermissionsRequest {
    #[serde(flatten)]
    pub page_request: PageRequest,
}

#[derive(Debug, Serialize, ToSchema)]
pub struct ListRolePermissionsResponse {
    pub permission_id: i64,
    pub object: String,
    pub action: String,
    pub name: Option<String>,
    pub description: Option<String>,
}

#[endpoint(
    tags("Role"),
)]
pub async fn page_role_permissions(
    role_id: PathParam<i64>,
    req: QueryParam<ListRolePermissionsRequest>
) -> ApiResult<PageResponse<ListRolePermissionsResponse>> {
    let db = utils::db::conn();
    
    let query = entity::role_permission::Entity::find()
        .filter(entity::role_permission::Column::RoleId.eq(role_id.into_inner()));
        
    let (offset, limit) = parse_page_request(req.page_request.clone());

    let total = query.clone().count(db).await.map_err(|e| {
        tracing::error!(error = ?e, "Failed to count role permission");
        ApiError::DbError(None)
    })?;

    let permissions: Vec<(entity::role_permission::Model, Option<entity::permission::Model>)> = query
        .order_by_asc(entity::permission::Column::PermissionId)
        .limit(limit)
        .offset(offset)
        .find_also_related(entity::permission::Entity)
        .all(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to query role permission");
            ApiError::DbError(None)
        })?;

    let resp = PageResponse {
        total,
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

    Ok(ApiOk(Some(resp)))
}