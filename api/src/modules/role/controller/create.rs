use salvo::{oapi::extract::JsonBody, prelude::*};
use serde::Deserialize;
use utils::response::{ApiOk, ApiResult};
use validator::Validate;

use crate::modules::role::service;

#[derive(Debug, Deserialize, Validate, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct CreateRoleRequest {
    pub code: String,
    pub reference_type: Option<String>,
    pub reference_id: Option<i64>,
    pub name: String,
    pub description: Option<String>,
}

/// Create a new role
#[endpoint(
    tags("Role"),
)]
pub async fn create_role(
    body: JsonBody<CreateRoleRequest>
) -> ApiResult<i64> {
    let new_role = service::create::create_role_by_request(body.into_inner()).await?;

    Ok(ApiOk(Some(new_role.role_id)))
}
