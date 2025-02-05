use salvo::{oapi::extract::PathParam, prelude::*};
use utils::response::{ApiOk, ApiResult};

use crate::modules::article::service;

/// Delete an article
#[endpoint(
    tags("Article"),
)]
pub async fn delete_article(
    article_id: PathParam<i64>,
) -> ApiResult<bool> {
    service::delete::delete_article_by_id(article_id.into_inner()).await?;

    Ok(ApiOk(Some(true)))
}