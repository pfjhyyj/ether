use salvo::{oapi::extract::JsonBody, prelude::*};
use entity::user;
use sea_orm::{ColumnTrait, EntityTrait, QueryFilter, ActiveModelTrait, Set};
use serde::{Deserialize, Serialize};

use utils::{identity::Identity, response::{ApiError, ApiOk, ApiResult}};
use validator::Validate;

#[derive(Debug, Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct GetCurrentUserResponse {
    pub user_id: i64,
    pub username: String,
    pub nickname: Option<String>,
    pub avatar: Option<String>,
}

/// Get current user
#[endpoint(
    tags("Current User"),
)]
pub async fn get_current_user(
    req: &mut Request
) -> ApiResult<GetCurrentUserResponse> {
    let id = req.extensions().get::<Identity>().unwrap();
    let user = get_user_by_user_id(id.sub).await?;

    let resp = GetCurrentUserResponse {
        user_id: user.user_id,
        username: user.username,
        nickname: user.nickname,
        avatar: user.avatar,
    };

    Ok(ApiOk(Some(resp)))
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

#[derive(Debug, Deserialize, Validate, ToSchema)]
pub struct UpdateCurrentUserRequest {
    pub nickname: Option<String>,
    pub avatar: Option<String>,
}

#[endpoint(
    tags("Current User"),
)]
pub async fn update_current_user(
    req: &mut Request,
    body: JsonBody<UpdateCurrentUserRequest>,
) -> ApiResult<GetCurrentUserResponse> {
    let id = req.extensions().get::<Identity>().unwrap();
    let user = update_user_by_user_id(id.sub, body.into_inner()).await?;

    let resp = GetCurrentUserResponse {
        user_id: user.user_id,
        username: user.username,
        nickname: user.nickname,
        avatar: user.avatar,
    };

    Ok(ApiOk(Some(resp)))
}

async fn update_user_by_user_id(user_id: i64, req: UpdateCurrentUserRequest) -> Result<user::Model, ApiError> {
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
