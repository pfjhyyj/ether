use salvo::{oapi::extract::JsonBody, prelude::*};
use entity::user;
use sea_orm::{ColumnTrait, EntityTrait, QueryFilter, ActiveModelTrait, Set};
use serde::Deserialize;
use utils::{identity::Identity, response::{ApiError, ApiOk, ApiResult}};
use validator::Validate;

#[derive(Debug, Deserialize, Validate, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct UpdatePasswordRequest {
    pub old_password: String,
    pub new_password: String,
    pub confirm_password: String,
}

/// Update current user password
#[endpoint(
    tags("Current User"),
)]
pub async fn update_password(
    req: &mut Request,
    body: JsonBody<UpdatePasswordRequest>,
) -> ApiResult<()> {
    let id = req.extensions().get::<Identity>().unwrap();
    let user = get_user_by_user_id(id.sub).await?;

    let is_valid = utils::hash::bcrypt_verify(&body.old_password, &user.password);
    if !is_valid {
        return Err(ApiError::RequestError(Some("Invalid old password".to_string())));
    }

    if body.new_password != body.confirm_password {
        return Err(ApiError::RequestError(Some("New password and confirm password do not match".to_string())));
    }

    let new_password = utils::hash::bcrypt(&body.new_password);

    let mut user: user::ActiveModel = user.into();
    user.password = Set(new_password);

    user.update(utils::db::conn()).await.map_err(|e| {
        tracing::error!(error = ?e, "Failed to update user password");
        ApiError::DbError(None)
    })?;

    Ok(ApiOk(None))
}

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