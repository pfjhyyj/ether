use salvo::prelude::*;
use salvo::oapi::{endpoint, extract::JsonBody, ToSchema};
use serde::{Deserialize, Serialize};
use utils::response::{ApiOk, ApiResult};
use validator::Validate;

use crate::modules::auth::service;

#[derive(Debug, Deserialize, ToSchema, Validate)]
pub struct LoginByUserNameRequest {
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

#[derive(Debug, Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct LoginByUserNameResponse {
    pub access_token: String,
    pub expire_time: i64,
}

/// Login by username and password
#[endpoint(
    tags("Auth"),
)]
pub async fn login_by_username(
    req: JsonBody<LoginByUserNameRequest>,
) -> ApiResult<LoginByUserNameResponse> {
    let resp = service::login::login_by_username(req.into_inner()).await?;
    Ok(ApiOk(Some(resp)))
}
