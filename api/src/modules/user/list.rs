use salvo::prelude::*;
use entity::user;
use sea_orm::{ColumnTrait, EntityTrait, PaginatorTrait, QueryFilter, QueryOrder, QuerySelect};
use serde::{Deserialize, Serialize};
use utils::{request::parse_page_and_size, response::{ApiError, ApiOk, ApiResult, PageResponse}};

#[derive(Debug, Deserialize, ToParameters)]
#[salvo(parameters(default_parameter_in = Query))]
pub struct PageUserRequest {
    pub page: Option<u64>,
    pub size: Option<u64>,
    pub username: Option<String>,
    pub nickname: Option<String>,
}

#[derive(Debug, Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct PageUserResponse {
    pub user_id: i64,
    pub username: String,
    pub nickname: Option<String>,
    pub avatar: Option<String>,
}

/// page users
#[endpoint(
    tags("User"),
)]
pub async fn page_user(
    req: PageUserRequest
) -> ApiResult<PageResponse<PageUserResponse>> {
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

    Ok(ApiOk(Some(resp)))
}