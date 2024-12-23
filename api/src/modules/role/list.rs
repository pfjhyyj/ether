use salvo::prelude::*;
use sea_orm::{ColumnTrait, EntityTrait, PaginatorTrait, QueryFilter, QueryOrder, QuerySelect};
use serde::{Deserialize, Serialize};
use utils::{request::parse_page_and_size, response::{ApiError, ApiOk, ApiResult, PageResponse}};

#[derive(Debug, Deserialize, ToParameters)]
#[salvo(parameters(default_parameter_in = Query))]
pub struct PageRoleRequest {
    pub page: Option<u64>,
    pub size: Option<u64>,
    pub name: Option<String>,
}

#[derive(Debug, Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct PageRoleResponse {
    pub role_id: i64,
    pub code: String,
    pub reference_type: Option<String>,
    pub reference_id: Option<i64>,
    pub name: String,
    pub description: Option<String>,
}

/// page roles
#[endpoint(
    tags("Role"),
)]
pub async fn page_role(
    req: PageRoleRequest
) -> ApiResult<PageResponse<PageRoleResponse>> {
    let db = utils::db::conn();
    let mut query = entity::role::Entity::find();

    if let Some(name) = &req.name {
        query = query.filter(entity::role::Column::Name.contains(name));
    }

    let (offset, limit) = parse_page_and_size(req.page, req.size);

    let total = query.clone().count(db).await.map_err(|e| {
        tracing::error!(error = ?e, "Failed to count role");
        ApiError::DbError(None)
    })?;

    let roles = query
        .order_by_asc(entity::role::Column::RoleId)
        .limit(limit)
        .offset(offset)
        .all(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to query role");
            ApiError::DbError(None)
        })?;

    let resp = PageResponse {
        total,
        page: offset / limit + 1,
        size: limit,
        data: roles.into_iter().map(|role| PageRoleResponse {
            role_id: role.role_id,
            code: role.code,
            reference_type: role.reference_type,
            reference_id: role.reference_id,
            name: role.name,
            description: role.description,
        }).collect(),
    };

    Ok(ApiOk(Some(resp)))
}