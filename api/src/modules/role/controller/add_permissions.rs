use salvo::{oapi::extract::{JsonBody, PathParam}, prelude::*};
use serde::Deserialize;
use utils::response::{ApiOk, ApiResult};
use validator::Validate;

use crate::modules::role::service;


#[derive(Debug, Deserialize, Validate, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct AddRolePermissionsRequest {
    pub permission_ids: Vec<i64>,
}

/// Add permissions to a role
#[endpoint(
    tags("Role"),
)]
pub async fn add_role_permissions(
    role_id: PathParam<i64>,
    body: JsonBody<AddRolePermissionsRequest>
) -> ApiResult<bool> {
    service::add_permissions::add_role_permissions_by_request(role_id.into_inner(), body.into_inner()).await?;
    
    Ok(ApiOk(Some(true)))
}
