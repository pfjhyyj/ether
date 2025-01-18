use sea_orm::{EntityTrait, Set, ActiveModelTrait};
use utils::response::ApiError;
use domain::entity::permission;

use crate::modules::permission::controller::update::UpdatePermissionRequest;

pub async fn update_permission_by_request(permission_id: i64, req: UpdatePermissionRequest) -> Result<bool, ApiError> {
    let db = utils::db::conn();
    let permission = permission::Entity::find_by_id(permission_id)
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find permission");
            ApiError::DbError(None)
        })?;

    if permission.is_none() {
        return Err(ApiError::RequestError(Some("Permission not found".to_string())));
    }

    let mut permission: permission::ActiveModel = permission.unwrap().into();
    permission.object = Set(req.object);
    permission.action = Set(req.action);
    permission.name = Set(req.name);
    permission.description = Set(req.description);

    permission.save(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to update permission");
            ApiError::DbError(None)
        })?;

    Ok(true)
}