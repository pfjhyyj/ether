use sea_orm::{ColumnTrait, EntityTrait, PaginatorTrait, QueryFilter, QueryOrder, QuerySelect};
use domain::entity::permission;
use utils::{request::parse_page_and_size, response::{ApiError, PageResponse}};

use crate::modules::permission::controller::list::{PagePermissionResponse, PagePermssionRequest};

pub async fn get_page_permission(req: PagePermssionRequest) -> Result<PageResponse<PagePermissionResponse>, ApiError> {
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

    Ok(resp)
}