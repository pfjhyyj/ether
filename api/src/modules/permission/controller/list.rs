use salvo::prelude::*;
use serde::{Deserialize, Serialize};
use utils::response::{ApiOk, ApiResult, PageResponse};

use crate::modules::permission::service;

#[derive(Debug, Deserialize, ToParameters)]
#[salvo(parameters(default_parameter_in = Query))]
pub struct PagePermssionRequest {
    pub page: Option<u64>,
    pub size: Option<u64>,
    pub object: Option<String>,
    pub action: Option<String>,
}

#[derive(Debug, Serialize, ToSchema)]
pub struct PagePermissionResponse {
    pub permission_id: i64,
    pub object: String,
    pub action: String,
    pub name: Option<String>,
    pub description: Option<String>,
}

/// page permissions
#[endpoint(
    tags("Permission"),
)]
pub async fn page_permission(
    req: PagePermssionRequest
) -> ApiResult<PageResponse<PagePermissionResponse>> {
    
    let resp = service::list::get_page_permission(req).await?;

    Ok(ApiOk(Some(resp)))
}