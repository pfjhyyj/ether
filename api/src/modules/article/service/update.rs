use domain::entity::article;
use sea_orm::{EntityTrait, Set, ActiveModelTrait, ColumnTrait, QueryFilter};
use utils::response::ApiError;

use crate::modules::article::controller::update::UpdateArticleRequest;

pub async fn update_article_by_request(article_id: i64, req: UpdateArticleRequest) -> Result<bool, ApiError> {
    let db = utils::db::conn();

    let article = article::Entity::find_by_id(article_id)
        .filter(article::Column::DeletedAt.is_null())
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find article");
            ApiError::DbError(None)
        })?;

    if article.is_none() {
        return Err(ApiError::RequestError(Some("Article not found".to_string())));
    }

    let mut article: article::ActiveModel = article.unwrap().into();
    article.title = Set(req.title);
    article.content = Set(req.content);
    article.category = Set(req.category);
    article.tag = Set(req.tag);
    article.status = Set(req.status);
    article.updated_at = Set(chrono::Utc::now().naive_utc());

    article.save(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to update article");
            ApiError::DbError(None)
        })?;

    Ok(true)
}