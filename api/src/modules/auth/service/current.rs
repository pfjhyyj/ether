use domain::entity::user;
use sea_orm::{ColumnTrait, EntityTrait, QueryFilter, ActiveModelTrait, Set};
use utils::response::ApiError;

use crate::modules::auth::controller::current::UpdateCurrentUserRequest;

pub async fn get_user_by_user_id(user_id: i64) -> Result<user::Model, ApiError> {
    let db = utils::db::conn();
    let user = user::Entity::find()
        .filter(user::Column::UserId.eq(user_id))
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to query user by user_id");
            ApiError::DbError(None)
        })?;

    user.ok_or(ApiError::RequestError(Some("User not found".to_string())))
}

pub async fn update_user_by_user_id(user_id: i64, req: UpdateCurrentUserRequest) -> Result<user::Model, ApiError> {
    let db = utils::db::conn();
    let user = user::Entity::find()
        .filter(user::Column::UserId.eq(user_id))
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to query user by user_id");
            ApiError::DbError(None)
        })?;

    let user = user.ok_or(ApiError::RequestError(Some("User not found".to_string())))?;

    let mut user: user::ActiveModel = user.into();
    if req.nickname.is_some() {
        user.nickname = Set(req.nickname.to_owned());
    }
    if req.avatar.is_some() {
        user.avatar = Set(req.avatar.to_owned());
    }

    let user = user.update(db).await.map_err(|e| {
        tracing::error!(error = ?e, "Failed to update user by user_id");
        ApiError::DbError(None)
    })?;
    Ok(user)
}
