use salvo::prelude::*;
use salvo::oapi::{endpoint, extract::JsonBody, ToSchema};
use serde::Deserialize;
use utils::response::{ApiOk, ApiResult};
use validator::Validate;

use crate::modules::user::service;



#[derive(Debug, Deserialize, Validate, ToSchema)]
pub struct CreateUserRequest {
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

/// Create a new user
#[endpoint(
    tags("User"),
)]
pub async fn create_user(
    body: JsonBody<CreateUserRequest>,
) -> ApiResult<i64> {
    let new_user = service::create::create_user_by_request(body.into_inner()).await?;

    Ok(ApiOk(Some(new_user.user_id)))
}
