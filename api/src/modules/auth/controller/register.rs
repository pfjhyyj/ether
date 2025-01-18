use salvo::prelude::*;
use salvo::oapi::{endpoint, extract::JsonBody, ToSchema};
use serde::Deserialize;
use utils::response::{ApiOk, ApiResult};
use validator::Validate;

use crate::modules::auth::service;

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

/// Register by username and password
#[endpoint(
    tags("Auth"),
)]
pub async fn register_by_username(
    req: JsonBody<RegisterByUsernameRequest>,
) -> ApiResult<i64> {
    let new_user = service::register::create_user_by_register_request(req.into_inner()).await?;

    Ok(ApiOk(Some(new_user.user_id)))
}
