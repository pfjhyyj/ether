use salvo::{oapi::extract::{JsonBody, PathParam}, prelude::*};
use serde::Deserialize;
use utils::response::{ApiError, ApiOk, ApiResult};
use sea_orm::{ColumnTrait, EntityTrait, QueryFilter};
use validator::Validate;


#[derive(Deserialize, Validate, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct RemoveRolePermissionsRequest {
    pub permission_ids: Vec<i64>,
}

/// Remove permissions from a role
#[endpoint(
    tags("Role"),
)]
pub async fn remove_role_permissions(
    role_id: PathParam<i64>,
    body: JsonBody<RemoveRolePermissionsRequest>
) -> ApiResult<bool> {
    let _ = remove_role_permissions_by_request(role_id.into_inner(), body.into_inner()).await?;
    Ok(ApiOk(Some(true)))
}

async fn remove_role_permissions_by_request(role_id: i64, req: RemoveRolePermissionsRequest) -> Result<bool, ApiError> {
    let db = utils::db::conn();

    let _ = entity::role_permission::Entity::delete_many()
        .filter(entity::role_permission::Column::RoleId.eq(role_id))
        .filter(entity::role_permission::Column::PermissionId.is_in(req.permission_ids.clone()))
        .exec(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to delete role permissions");
            ApiError::DbError(None)
        })?;

    Ok(true)
}