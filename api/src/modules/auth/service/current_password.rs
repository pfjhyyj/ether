use domain::entity::user;
use sea_orm::{ColumnTrait, EntityTrait, QueryFilter, ActiveModelTrait, Set};
use utils::response::{ApiError, ApiOk, ApiResult};


async fn get_user_by_user_id(user_id: i64) -> Result<user::Model, ApiError> {
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

pub async fn update_password(
    user_id: i64,
    old_password: &str,
    new_password: &str,
) -> ApiResult<bool> {
    let user = get_user_by_user_id(user_id).await?;

    let is_valid = utils::hash::bcrypt_verify(old_password, &user.password);
    if !is_valid {
        return Err(ApiError::RequestError(Some("Invalid old password".to_string())));
    }

    let new_password = utils::hash::bcrypt(new_password);

    let mut user: user::ActiveModel = user.into();
    user.password = Set(new_password);

    user.update(utils::db::conn()).await.map_err(|e| {
        tracing::error!(error = ?e, "Failed to update user password");
        ApiError::DbError(None)
    })?;

    Ok(ApiOk(Some(true)))
}