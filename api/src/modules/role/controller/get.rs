use salvo::{oapi::extract::PathParam, prelude::*};
use serde::Serialize;
use utils::response::{ApiOk, ApiResult};

use crate::modules::role::service;

#[derive(Debug, Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct GetRoleDetailResponse {
    pub role_id: i64,
    pub code: String,
    pub reference_type: Option<String>,
    pub reference_id: Option<i64>,
    pub name: String,
    pub description: Option<String>,
}

/// Get a role
#[endpoint(
    tags("Role"),
)]
pub async fn get_role(
    role_id: PathParam<i64>,
) -> ApiResult<GetRoleDetailResponse> {
    let role = service::get::get_role_by_id(role_id.into_inner()).await?;
    let role = GetRoleDetailResponse {
        role_id: role.role_id,
        code: role.code,
        reference_type: role.reference_type,
        reference_id: role.reference_id,
        name: role.name,
        description: role.description,
    };
    Ok(ApiOk(Some(role)))
}
