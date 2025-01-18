use domain::entity::user;
use sea_orm::{ColumnTrait, EntityTrait, PaginatorTrait, QueryFilter, QueryOrder, QuerySelect};
use utils::{request::parse_page_and_size, response::{ApiError, PageResponse}};

use crate::modules::user::controller::list::{PageUserRequest, PageUserResponse};

pub async fn get_page_user(
    req: PageUserRequest
) -> Result<PageResponse<PageUserResponse>, ApiError> {
    let db = utils::db::conn();
    let mut query = user::Entity::find();

    if let Some(username) = &req.username {
        query = query.filter(user::Column::Username.contains(username));
    }

    if let Some(nickname) = &req.nickname {
        query = query.filter(user::Column::Nickname.contains(nickname));
    }

    let (offset, limit) = parse_page_and_size(req.page, req.size);

    let total = query.clone().count(db).await.map_err(|e| {
        tracing::error!(error = ?e, "Failed to count user");
        ApiError::DbError(None)
    })?;

    let users = query
        .order_by_asc(user::Column::UserId)
        .limit(limit)
        .offset(offset)
        .all(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to query user");
            ApiError::DbError(None)
        })?;

    let resp = PageResponse {
        total,
        page: offset / limit + 1,
        size: limit,
        data: users.into_iter().map(|user| PageUserResponse {
            user_id: user.user_id,
            username: user.username,
            nickname: user.nickname,
            avatar: user.avatar,
        }).collect(),
    };

    Ok(resp)
}