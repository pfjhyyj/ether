use entity::user;
use salvo::prelude::*;
use salvo::oapi::{endpoint, extract::{JsonBody, PathParam}, ToSchema};
use sea_orm::{EntityTrait, Set, ActiveModelTrait};
use serde::Deserialize;
use utils::response::{ApiError, ApiOk, ApiResult};
use validator::Validate;



#[derive(Debug, Deserialize, Validate, ToSchema)]
pub struct UpdateUserRequest {
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
    pub email: Option<String>,
    pub nickname: Option<String>,
}

/// Update a user
#[endpoint(
    tags("User"),
)]
pub async fn update_user(
    user_id: PathParam<i64>,
    body: JsonBody<UpdateUserRequest>,
) -> ApiResult<bool> {
    let _ = update_user_by_request(user_id.into_inner(), body.into_inner()).await?;

    Ok(ApiOk(Some(true)))
}

async fn update_user_by_request(user_id: i64, req: UpdateUserRequest) -> Result<bool, ApiError> {
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