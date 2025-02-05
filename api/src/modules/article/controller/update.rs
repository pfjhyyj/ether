use salvo::{oapi::extract::{JsonBody, PathParam}, prelude::*};
use serde::Deserialize;
use utils::response::{ApiOk, ApiResult};

use crate::modules::article::service;

#[derive(Deserialize, ToSchema)]
pub struct UpdateArticleRequest {
    pub title: String,
    pub content: Option<String>,
    pub category: Option<String>,
    pub tag: Option<String>,
    pub status: i32,
}

/// Update an article
#[endpoint(
    tags("Article"),
)]
pub async fn update_article(
    article_id: PathParam<i64>,
    body: JsonBody<UpdateArticleRequest>,
) -> ApiResult<bool> {
    service::update::update_article_by_request(article_id.into_inner(), body.into_inner()).await?;

    Ok(ApiOk(Some(true)))
}