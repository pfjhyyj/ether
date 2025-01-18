use salvo::{oapi::extract::{JsonBody, PathParam}, prelude::*};
use serde::Deserialize;
use utils::response::{ApiOk, ApiResult};
use validator::Validate;

use crate::modules::permission::service;



#[derive(Debug, Deserialize, Validate, ToSchema)]
pub struct UpdatePermissionRequest {
    pub object: String,
    pub action: String,
    pub name: Option<String>,
    pub description: Option<String>,
}

/// Update a permission
#[endpoint(
    tags("Permission"),
)]
pub async fn update_permission(
    permission_id: PathParam<i64>,
    body: JsonBody<UpdatePermissionRequest>,
) -> ApiResult<bool> {
    let _ = service::update::update_permission_by_request(permission_id.into_inner(), body.into_inner()).await?;

    Ok(ApiOk(Some(true)))
}
