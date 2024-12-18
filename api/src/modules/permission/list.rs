use salvo::{oapi::extract::QueryParam, prelude::*};
use sea_orm::{ColumnTrait, EntityTrait, PaginatorTrait, QueryFilter, QueryOrder, QuerySelect};
use serde::{Deserialize, Serialize};
use utils::{request::{parse_page_request, PageRequest}, response::{ApiError, ApiOk, ApiResult, PageResponse}};

#[derive(Debug, Deserialize, ToSchema)]
pub struct PagePermssionRequest {
    #[serde(flatten)]
    pub page_request: PageRequest,
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

#[endpoint(
    tags("Permission"),
)]
pub async fn page_permission(
    req: QueryParam<PagePermssionRequest>
) -> ApiResult<PageResponse<PagePermissionResponse>> {
    let db = utils::db::conn();
    let mut query = entity::permission::Entity::find();

    if let Some(object) = &req.object {
        query = query.filter(entity::permission::Column::Object.contains(object));
    }

    if let Some(action) = &req.action {
        query = query.filter(entity::permission::Column::Action.contains(action));
    }

    let (offset, limit) = parse_page_request(req.into_inner().page_request);

    let total = query.clone().count(db).await.map_err(|e| {
        tracing::error!(error = ?e, "Failed to count permission");
        ApiError::DbError(None)
    })?;

    let permissions = query
        .order_by_asc(entity::permission::Column::PermissionId)
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