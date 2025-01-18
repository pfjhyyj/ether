use salvo::prelude::*;
use sea_orm::{ColumnTrait, EntityTrait, PaginatorTrait, QueryFilter, QueryOrder, QuerySelect};
use serde::{Deserialize, Serialize};
use domain::entity::permission;
use utils::{request::parse_page_and_size, response::{ApiError, ApiOk, ApiResult, PageResponse}};

#[derive(Debug, Deserialize, ToParameters)]
#[salvo(parameters(default_parameter_in = Query))]
pub struct PagePermssionRequest {
    pub page: Option<u64>,
    pub size: Option<u64>,
    pub object: Option<String>,
    pub action: Option<String>,
}

#[derive(Debug, Serialize, ToSchema)]
pub struct PagePermissionResponse {
    pub permission_id: i64,
    pub object: String,
    pub action: String,
    pub name: Option<String>,
    pub description: Option<String>,
}

/// page permissions
#[endpoint(
    tags("Permission"),
)]
pub async fn page_permission(
    req: PagePermssionRequest
) -> ApiResult<PageResponse<PagePermissionResponse>> {
    let db = utils::db::conn();
    let mut query = permission::Entity::find();

    if let Some(object) = &req.object {
        query = query.filter(permission::Column::Object.contains(object));
    }

    if let Some(action) = &req.action {
        query = query.filter(permission::Column::Action.contains(action));
    }

    let (offset, limit) = parse_page_and_size(req.page, req.size);

    let total = query.clone().count(db).await.map_err(|e| {
        tracing::error!(error = ?e, "Failed to count permission");
        ApiError::DbError(None)
    })?;

    let permissions = query
        .order_by_asc(permission::Column::PermissionId)
        .limit(limit)
        .offset(offset)
        .all(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to query permission");
            ApiError::DbError(None)
        })?;

    let resp = PageResponse {
        total,
        page: offset / limit + 1,
        size: limit,
        data: permissions.into_iter().map(|permission| PagePermissionResponse {
            permission_id: permission.permission_id,
            object: permission.object,
            action: permission.action,
            name: permission.name,
            description: permission.description,
        }).collect(),
    };

    Ok(ApiOk(Some(resp)))
}