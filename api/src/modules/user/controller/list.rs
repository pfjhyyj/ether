use salvo::prelude::*;
use serde::{Deserialize, Serialize};
use utils::response::{ApiOk, ApiResult, PageResponse};

use crate::modules::user::service;

#[derive(Debug, Deserialize, ToParameters)]
#[salvo(parameters(default_parameter_in = Query))]
pub struct PageUserRequest {
    pub page: Option<u64>,
    pub size: Option<u64>,
    pub username: Option<String>,
    pub nickname: Option<String>,
}

#[derive(Debug, Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct PageUserResponse {
    pub user_id: i64,
    pub username: String,
    pub nickname: Option<String>,
    pub avatar: Option<String>,
}

/// page users
#[endpoint(
    tags("User"),
)]
pub async fn page_user(
    req: PageUserRequest
) -> ApiResult<PageResponse<PageUserResponse>> {
    let resp = service::list::get_page_user(req).await?;

    Ok(ApiOk(Some(resp)))
}