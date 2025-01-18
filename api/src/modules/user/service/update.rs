use domain::entity::user;
use sea_orm::{EntityTrait, Set, ActiveModelTrait};
use utils::response::ApiError;

use crate::modules::user::controller::update::UpdateUserRequest;

pub async fn update_user_by_request(user_id: i64, req: UpdateUserRequest) -> Result<bool, ApiError> {
    let db = utils::db::conn();
    let user = user::Entity::find_by_id(user_id)
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find user");
            ApiError::DbError(None)
        })?;

    if user.is_none() {
        return Err(ApiError::RequestError(Some("User not found".to_string())));
    }

    let mut user: user::ActiveModel = user.unwrap().into();
    user.username = Set(req.username);
    user.password = Set(utils::hash::bcrypt(&req.password));
    user.email = Set(req.email);
    user.nickname = Set(req.nickname);

    user.save(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to update user");
            ApiError::DbError(None)
        })?;

    Ok(true)
}