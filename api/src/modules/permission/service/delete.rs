use sea_orm::{EntityTrait, ModelTrait};
use domain::entity::permission;
use utils::response::ApiError;

pub async fn delete_permission_by_id(permission_id: i64) -> Result<bool, ApiError> {
    let db = utils::db::conn();
    let permission = permission::Entity::find_by_id(permission_id)
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find permission");
            ApiError::DbError(None)
        })?;
    
    if let Some(permission) = permission {
        permission.delete(db)
            .await
            .map_err(|e| {
                tracing::error!(error = ?e, "Failed to delete permission");
                ApiError::DbError(None)
            })?;
    }
    Ok(true)
}