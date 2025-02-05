use salvo::{oapi::extract::JsonBody, prelude::*};
use serde::Deserialize;
use utils::response::{ApiOk, ApiResult};
use validator::Validate;

use crate::modules::article::service;



#[derive(Deserialize, Validate, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct CreateArticleRequest {
    pub title: String,
    pub content: Option<String>,
    pub category: Option<String>,
    pub tag: Option<String>,
    pub status: i32,
}

/// Create a new article
#[endpoint(
    tags("Article"),
)]
pub async fn create_article(
    body: JsonBody<CreateArticleRequest>,
) -> ApiResult<i64> {
    let new_article = service::create::create_article_by_request(body.into_inner()).await?;

    Ok(ApiOk(Some(new_article.article_id)))
}