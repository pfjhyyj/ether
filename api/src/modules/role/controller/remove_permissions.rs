use salvo::{oapi::extract::{JsonBody, PathParam}, prelude::*};
use serde::Deserialize;
use utils::response::{ApiOk, ApiResult};
use validator::Validate;

use crate::modules::role::service;


#[derive(Deserialize, Validate, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct RemoveRolePermissionsRequest {
    pub permission_ids: Vec<i64>,
}

/// Remove permissions from a role
#[endpoint(
    tags("Role"),
)]
pub async fn remove_role_permissions(
    role_id: PathParam<i64>,
    body: JsonBody<RemoveRolePermissionsRequest>
) -> ApiResult<bool> {
    service::remove_permissions::remove_role_permissions_by_request(role_id.into_inner(), body.into_inner()).await?;
    Ok(ApiOk(Some(true)))
}
