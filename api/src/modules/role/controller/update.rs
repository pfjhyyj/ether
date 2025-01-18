use salvo::{oapi::extract::{JsonBody, PathParam}, prelude::*};
use serde::Deserialize;
use utils::response::{ApiOk, ApiResult};
use validator::Validate;

use crate::modules::role::service;

#[derive(Debug, Deserialize, Validate, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct UpdateRoleRequest {
    pub code: String,
    pub reference_type: Option<String>,
    pub reference_id: Option<i64>,
    pub name: String,
    pub description: Option<String>,
}

/// Update a role
#[endpoint(
    tags("Role"),
)]
pub async fn update_role(
    role_id: PathParam<i64>,
    body: JsonBody<UpdateRoleRequest>,
) -> ApiResult<bool> {
    service::update::update_role_by_request(role_id.into_inner(), body.into_inner()).await?;

    Ok(ApiOk(Some(true)))
}
