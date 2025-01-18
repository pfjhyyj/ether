use salvo::{oapi::extract::PathParam, prelude::*};
use serde::{Deserialize, Serialize};
use utils::response::{ApiOk, ApiResult, PageResponse};

use crate::modules::role::service;


#[derive(Debug, Deserialize, ToParameters)]
#[salvo(parameters(default_parameter_in = Query))]
pub struct ListRolePermissionsRequest {
    pub page: Option<u64>,
    pub size: Option<u64>,
}

#[derive(Debug, Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct ListRolePermissionsResponse {
    pub permission_id: i64,
    pub object: String,
    pub action: String,
    pub name: Option<String>,
    pub description: Option<String>,
}

/// page role permissions
#[endpoint(
    tags("Role"),
)]
pub async fn page_role_permissions(
    role_id: PathParam<i64>,
    req: ListRolePermissionsRequest
) -> ApiResult<PageResponse<ListRolePermissionsResponse>> {

    let resp = service::list_permissions::get_page_role_permissions(role_id.into_inner(), req).await?;

    Ok(ApiOk(Some(resp)))
}