use salvo::prelude::*;
use domain::entity::user;
use salvo::oapi::extract::PathParam;
use sea_orm::EntityTrait;
use serde::Serialize;
use utils::response::{ApiError, ApiOk, ApiResult};

#[derive(Debug, Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct GetUserDetailResponse {
    pub user_id: i64,
    pub username: String,
    pub email: Option<String>,
    pub nickname: Option<String>,
    pub avatar: Option<String>,
    pub created_at: chrono::NaiveDateTime,
    pub updated_at: chrono::NaiveDateTime,
}

/// Get a user
#[endpoint(
    tags("User"),
)]
pub async fn get_user(
    user_id: PathParam<i64>,
) -> ApiResult<GetUserDetailResponse> {
    let user = get_user_by_id(user_id.into_inner()).await?;
    let user = GetUserDetailResponse {
        user_id: user.user_id,
        username: user.username,
        email: user.email,
        nickname: user.nickname,
        avatar: user.avatar,
        created_at: user.created_at.naive_local(),
        updated_at: user.updated_at.naive_local(),
    };

    Ok(ApiOk(Some(user)))
}

async fn get_user_by_id(id: i64) -> Result<user::Model, ApiError> {
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