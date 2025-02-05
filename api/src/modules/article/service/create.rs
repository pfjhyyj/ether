use domain::entity::article;
use sea_orm::{Set, ActiveModelTrait};
use utils::response::ApiError;

use crate::modules::article::controller::create::CreateArticleRequest;



pub async fn create_article_by_request(req: CreateArticleRequest) -> Result<article::Model, ApiError> {
    let db = utils::db::conn();

    let new_article = article::ActiveModel {
        title: Set(req.title),
        content: Set(req.content),
        category: Set(req.category),
        tag: Set(req.tag),
        status: Set(req.status),
        ..Default::default()
    }.insert(db);

    let new_article = new_article.await.map_err(|e| {
        tracing::error!(error = ?e, "Failed to insert new article");
        ApiError::DbError(None)
    })?;

    Ok(new_article)
}