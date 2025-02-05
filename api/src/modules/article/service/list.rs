use domain::entity::article;
use sea_orm::{EntityTrait, PaginatorTrait, QueryOrder, QuerySelect, ColumnTrait, QueryFilter};
use utils::{request::parse_page_and_size, response::{ApiError, PageResponse}};

use crate::modules::article::controller::list::{PageArticleRequest, PageArticleResponse};

pub async fn get_page_article(
    req: PageArticleRequest
) -> Result<PageResponse<PageArticleResponse>, ApiError> {
    let db = utils::db::conn();
    let mut query = article::Entity::find();
    query = query.filter(article::Column::DeletedAt.is_null());

    let (offset, limit) = parse_page_and_size(req.page, req.size);

    let total = query.clone().count(db).await.map_err(|e| {
        tracing::error!(error = ?e, "Failed to count article");
        ApiError::DbError(None)
    })?;

    let articles = query
        .order_by_asc(article::Column::ArticleId)
        .limit(limit)
        .offset(offset)
        .all(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to query article");
            ApiError::DbError(None)
        })?;

    let resp = PageResponse {
        total,
        page: offset / limit + 1,
        size: limit,
        data: articles.into_iter().map(|article| PageArticleResponse {
            article_id: article.article_id,
            title: article.title,
            content: article.content,
            category: article.category,
            tag: article.tag,
            status: article.status,
            created_at: article.created_at.and_utc().timestamp(),
            updated_at: article.updated_at.and_utc().timestamp(),
        }).collect(),
    };

    Ok(resp)
}