use domain::entity::article;
use sea_orm::{EntityTrait, QueryFilter, ColumnTrait};
use utils::response::ApiError;

pub async fn get_article_by_id(article_id: i64) -> Result<article::Model, ApiError> {
    let db = utils::db::conn();

    let article = article::Entity::find_by_id(article_id)
        .filter(article::Column::DeletedAt.is_null())
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find article");
            ApiError::DbError(None)
        })?;
    
    if let Some(article) = article {
        Ok(article)
    } else {
        Err(ApiError::RequestError(Some("Article not found".to_string())))
    }
}