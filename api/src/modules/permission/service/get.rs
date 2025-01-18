use sea_orm::EntityTrait;
use domain::entity::permission;
use utils::response::ApiError;

pub async fn get_permission_by_id(id: i64) -> Result<permission::Model, ApiError> {
    let db = utils::db::conn();
    let permission = permission::Entity::find_by_id(id)
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find permission");
            utils::response::ApiError::DbError(None)
        })?;

    if let Some(permission) = permission {
        Ok(permission)
    } else {
        Err(utils::response::ApiError::RequestError(Some("Permission not found".to_string())))
    }
}