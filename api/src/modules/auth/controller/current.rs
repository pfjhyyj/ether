use salvo::{oapi::extract::JsonBody, prelude::*};
use serde::{Deserialize, Serialize};

use utils::{identity::Identity, response::{ApiOk, ApiResult}};
use validator::Validate;

use crate::modules::auth::service;

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
    let user = service::current::get_user_by_user_id(id.sub).await?;

    let resp = GetCurrentUserResponse {
        user_id: user.user_id,
        username: user.username,
        nickname: user.nickname,
        avatar: user.avatar,
    };

    Ok(ApiOk(Some(resp)))
}

#[derive(Debug, Deserialize, Validate, ToSchema)]
pub struct UpdateCurrentUserRequest {
    pub nickname: Option<String>,
    pub avatar: Option<String>,
}

/// Update current user
#[endpoint(
    tags("Current User"),
)]
pub async fn update_current_user(
    req: &mut Request,
    body: JsonBody<UpdateCurrentUserRequest>,
) -> ApiResult<GetCurrentUserResponse> {
    let id = req.extensions().get::<Identity>().unwrap();
    let user = service::current::update_user_by_user_id(id.sub, body.into_inner()).await?;

    let resp = GetCurrentUserResponse {
        user_id: user.user_id,
        username: user.username,
        nickname: user.nickname,
        avatar: user.avatar,
    };

    Ok(ApiOk(Some(resp)))
}
