use salvo::{oapi::extract::JsonBody, prelude::*};
use serde::Deserialize;
use utils::{identity::Identity, response::{ApiError, ApiOk, ApiResult}};
use validator::Validate;

use crate::modules::auth::service;

#[derive(Debug, Deserialize, Validate, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct UpdatePasswordRequest {
    pub old_password: String,
    pub new_password: String,
    pub confirm_password: String,
}

/// Update current user password
#[endpoint(
    tags("Current User"),
)]
pub async fn update_password(
    req: &mut Request,
    body: JsonBody<UpdatePasswordRequest>,
) -> ApiResult<()> {
    let id = req.extensions().get::<Identity>().unwrap();
    if body.new_password != body.confirm_password {
        return Err(ApiError::RequestError(Some("New password and confirm password do not match".to_string())));
    }

    service::current_password::update_password(id.sub, &body.old_password, &body.new_password).await?;

    Ok(ApiOk(None))
}
