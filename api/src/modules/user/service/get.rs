use domain::entity::user;
use sea_orm::EntityTrait;
use utils::response::ApiError;

pub async fn get_user_by_id(id: i64) -> Result<user::Model, ApiError> {
    let db = utils::db::conn();
    let user = user::Entity::find_by_id(id)
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find user");
            ApiError::DbError(None)
        })?;

    if let Some(user) = user {
        Ok(user)
    } else {
        Err(ApiError::RequestError(Some("User not found".to_string())))
    }
}