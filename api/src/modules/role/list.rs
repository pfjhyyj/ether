use salvo::{oapi::extract::QueryParam, prelude::*};
use sea_orm::{ColumnTrait, EntityTrait, PaginatorTrait, QueryFilter, QueryOrder, QuerySelect};
use serde::{Deserialize, Serialize};
use utils::{request::{parse_page_request, PageRequest}, response::{ApiError, ApiOk, ApiResult, PageResponse}};

#[derive(Debug, Deserialize, ToSchema)]
pub struct PageRoleRequest {
    #[serde(flatten)]
    pub page_request: PageRequest,
    pub name: Option<String>,
}

#[derive(Debug, Serialize, ToSchema)]
pub struct PageRoleResponse {
    pub role_id: i64,
    pub code: String,
    pub reference_type: Option<String>,
    pub reference_id: Option<i64>,
    pub name: String,
    pub description: Option<String>,
}

#[endpoint(
    tags("Role"),
)]
pub async fn page_role(
    req: QueryParam<PageRoleRequest>
) -> ApiResult<PageResponse<PageRoleResponse>> {
    let db = utils::db::conn();
    let mut query = entity::role::Entity::find();

    if let Some(name) = &req.name {
        query = query.filter(entity::role::Column::Name.contains(name));
    }

    let (offset, limit) = parse_page_request(req.page_request.clone());

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