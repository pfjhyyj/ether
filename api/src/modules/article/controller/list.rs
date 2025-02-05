use salvo::prelude::*;
use serde::{Deserialize, Serialize};
use utils::response::{ApiOk, ApiResult, PageResponse};

use crate::modules::article::service;

#[derive(Deserialize, ToParameters)]
#[salvo(parameters(default_parameter_in = Query))]
pub struct PageArticleRequest {
    pub page: Option<u64>,
    pub size: Option<u64>,
}

#[derive(Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct PageArticleResponse {
    pub article_id: i64,
    pub title: String,
    pub content: Option<String>,
    pub category: Option<String>,
    pub tag: Option<String>,
    pub status: i32,
    pub created_at: i64,
    pub updated_at: i64,
}

/// page articles
#[endpoint(
    tags("Article"),
)]
pub async fn page_article(
    req: PageArticleRequest
) -> ApiResult<PageResponse<PageArticleResponse>> {
    let resp = service::list::get_page_article(req).await?;

    Ok(ApiOk(Some(resp)))
}