use sea_orm::EntityTrait;
use domain::entity::role;
use utils::response::ApiError;

pub async fn get_role_by_id(id: i64) -> Result<role::Model, ApiError> {
    let db = utils::db::conn();
    let role = role::Entity::find_by_id(id)
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find role");
            utils::response::ApiError::DbError(None)
        })?;

    if let Some(role) = role {
        Ok(role)
    } else {
        Err(utils::response::ApiError::RequestError(Some("Role not found".to_string())))
    }
}