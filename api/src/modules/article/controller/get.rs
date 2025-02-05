use salvo::{oapi::extract::PathParam, prelude::*};
use serde::Serialize;
use utils::response::{ApiOk, ApiResult};

use crate::modules::article::service;

#[derive(Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct GetArticleResponse {
    pub article_id: i64,
    pub title: String,
    pub content: Option<String>,
    pub category: Option<String>,
    pub tag: Option<String>,
    pub status: i32,
}

/// Get an article
#[endpoint(
    tags("Article"),
)]
pub async fn get_article(
    article_id: PathParam<i64>,
) -> ApiResult<GetArticleResponse> {
    let article = service::get::get_article_by_id(article_id.into_inner()).await?;
    let article = GetArticleResponse {
        article_id: article.article_id,
        title: article.title,
        content: article.content,
        category: article.category,
        tag: article.tag,
        status: article.status,
    };
    Ok(ApiOk(Some(article)))
}