use entity::user;
use salvo::prelude::*;
use salvo::oapi::{endpoint, extract::JsonBody, ToSchema};
use sea_orm::{ActiveModelTrait, ColumnTrait, EntityTrait, QueryFilter, Set};
use serde::Deserialize;
use utils::response::{ApiError, ApiOk, ApiResult};
use validator::Validate;

#[derive(Debug, Deserialize, Validate, ToSchema)]
pub struct RegisterByUsernameRequest {
    #[validate(length(
        min = 6,
        max = 50,
        message = "Username must be between 6 and 50 characters"
    ))]
    pub username: String,
    #[validate(length(
        min = 6,
        max = 50,
        message = "Password must be between 6 and 50 characters"
    ))]
    pub password: String,
}

#[endpoint(
    tags("Auth"),
)]
pub async fn register_by_username(
    req: JsonBody<RegisterByUsernameRequest>,
) -> ApiResult<i64> {
    let new_user = create_user_by_register_request(req.into_inner()).await?;

    Ok(ApiOk(Some(new_user.user_id)))
}

async fn create_user_by_register_request(req: RegisterByUsernameRequest) -> Result<user::Model, ApiError> {
    let db = utils::db::conn();
    let user = user::Entity::find()
        .filter(user::Column::Username.eq(&req.username))
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to query user by username");
            ApiError::DbError(None)
        })?;

    if user.is_some() {
        return Err(ApiError::RequestError(Some("Username already exists".to_string())));
    }

    let password = utils::hash::bcrypt(&req.password);

    let new_user = user::ActiveModel {
        username: Set(req.username),
        password: Set(password),
        ..Default::default()
    };

    let new_user = new_user.insert(db).await.map_err(|e| {
        tracing::error!(error = ?e, "Failed to insert new user");
        ApiError::DbError(None)
    })?;

    Ok(new_user)
}