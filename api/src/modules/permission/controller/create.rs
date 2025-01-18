use salvo::{oapi::extract::JsonBody, prelude::*};
use serde::Deserialize;
use utils::response::{ApiOk, ApiResult};
use validator::Validate;

use crate::modules::permission::service;


#[derive(Debug, Deserialize, Validate, ToSchema)]
pub struct CreatePermissionRequest {
    pub object: String,
    pub action: String,
    pub name: Option<String>,
    pub description: Option<String>,
}

/// Create a new permission
#[endpoint(
    tags("Permission"),
)]
pub async fn create_permission(
    body: JsonBody<CreatePermissionRequest>,
) -> ApiResult<i64> {
    let new_permission = service::create::create_permission_by_request(body.into_inner()).await?;

    Ok(ApiOk(Some(new_permission.permission_id)))
}