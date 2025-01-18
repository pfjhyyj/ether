use utils::response::ApiError;
use sea_orm::{ColumnTrait, EntityTrait, QueryFilter};
use domain::entity::role_permission;

use crate::modules::role::controller::remove_permissions::RemoveRolePermissionsRequest;

pub async fn remove_role_permissions_by_request(role_id: i64, req: RemoveRolePermissionsRequest) -> Result<bool, ApiError> {
    let db = utils::db::conn();

    let _ = role_permission::Entity::delete_many()
        .filter(role_permission::Column::RoleId.eq(role_id))
        .filter(role_permission::Column::PermissionId.is_in(req.permission_ids.clone()))
        .exec(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to delete role permissions");
            ApiError::DbError(None)
        })?;

    Ok(true)
}