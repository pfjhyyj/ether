use salvo::prelude::*;
use serde::{Deserialize, Serialize};
use utils::response::{ApiOk, ApiResult, PageResponse};

use crate::modules::role::service;

#[derive(Debug, Deserialize, ToParameters)]
#[salvo(parameters(default_parameter_in = Query))]
pub struct PageRoleRequest {
    pub page: Option<u64>,
    pub size: Option<u64>,
    pub name: Option<String>,
}

#[derive(Debug, Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct PageRoleResponse {
    pub role_id: i64,
    pub code: String,
    pub reference_type: Option<String>,
    pub reference_id: Option<i64>,
    pub name: String,
    pub description: Option<String>,
}

/// page roles
#[endpoint(
    tags("Role"),
)]
pub async fn page_role(
    req: PageRoleRequest
) -> ApiResult<PageResponse<PageRoleResponse>> {
    let resp = service::list::get_page_role(req).await?;

    Ok(ApiOk(Some(resp)))
}