use domain::entity::role;
use sea_orm::{ColumnTrait, EntityTrait, PaginatorTrait, QueryFilter, QueryOrder, QuerySelect};
use utils::{request::parse_page_and_size, response::{ApiError, PageResponse}};

use crate::modules::role::controller::list::{PageRoleRequest, PageRoleResponse};

pub async fn get_page_role(
    req: PageRoleRequest
) -> Result<PageResponse<PageRoleResponse>, ApiError> {
    let db = utils::db::conn();
    let mut query = role::Entity::find();

    if let Some(name) = &req.name {
        query = query.filter(role::Column::Name.contains(name));
    }

    let (offset, limit) = parse_page_and_size(req.page, req.size);

    let total = query.clone().count(db).await.map_err(|e| {
        tracing::error!(error = ?e, "Failed to count role");
        ApiError::DbError(None)
    })?;

    let roles = query
        .order_by_asc(role::Column::RoleId)
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

    Ok(resp)
}