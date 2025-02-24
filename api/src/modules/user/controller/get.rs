use salvo::prelude::*;
use salvo::oapi::extract::PathParam;
use serde::Serialize;
use utils::response::{ApiOk, ApiResult};

use crate::modules::user::service;

#[derive(Debug, Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct GetUserDetailResponse {
    pub user_id: i64,
    pub username: String,
    pub email: Option<String>,
    pub nickname: Option<String>,
    pub avatar: Option<String>,
}

/// Get a user
#[endpoint(
    tags("User"),
)]
pub async fn get_user(
    user_id: PathParam<i64>,
) -> ApiResult<GetUserDetailResponse> {
    let user = service::get::get_user_by_id(user_id.into_inner()).await?;
    let user = GetUserDetailResponse {
        user_id: user.user_id,
        username: user.username,
        email: user.email,
        nickname: user.nickname,
        avatar: user.avatar,
    };

    Ok(ApiOk(Some(user)))
}
