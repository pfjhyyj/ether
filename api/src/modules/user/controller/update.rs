use salvo::prelude::*;
use salvo::oapi::{endpoint, extract::{JsonBody, PathParam}, ToSchema};
use serde::Deserialize;
use utils::response::{ApiOk, ApiResult};
use validator::Validate;

use crate::modules::user::service;



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
    service::update::update_user_by_request(user_id.into_inner(), body.into_inner()).await?;

    Ok(ApiOk(Some(true)))
}
