use domain::entity::article;
use sea_orm::{EntityTrait, Set, QueryFilter, ColumnTrait};
use utils::response::ApiError;

pub async fn delete_article_by_id(article_id: i64) -> Result<bool, ApiError> {
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
    article.deleted_at = Set(Some(chrono::Utc::now().naive_utc()));

    
    Ok(true)
}