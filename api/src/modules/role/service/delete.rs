use sea_orm::{EntityTrait, ModelTrait};
use domain::entity::role;
use utils::response::ApiError;

pub async fn delete_role_by_id(role_id: i64) -> Result<bool, ApiError> {
    let db = utils::db::conn();
    let role = role::Entity::find_by_id(role_id)
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find role");
            ApiError::DbError(None)
        })?;
    
    if let Some(role) = role {
        role.delete(db)
            .await
            .map_err(|e| {
                tracing::error!(error = ?e, "Failed to delete role");
                ApiError::DbError(None)
            })?;
    }
    Ok(true)
}