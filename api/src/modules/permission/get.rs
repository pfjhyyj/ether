use salvo::{oapi::extract::PathParam, prelude::*};
use sea_orm::EntityTrait;
use serde::Serialize;
use utils::response::{ApiError, ApiOk, ApiResult};

#[derive(Debug, Serialize, ToSchema)]
pub struct GetPermssionDetailResponse {
    pub permission_id: i64,
    pub object: String,
    pub action: String,
    pub name: Option<String>,
    pub description: Option<String>,
}

#[endpoint(
    tags("Permission"),
)]
pub async fn get_permission(
    permission_id: PathParam<i64>,
) -> ApiResult<GetPermssionDetailResponse> {
    let permission = get_permission_by_id(permission_id.into_inner()).await?;
    let permission = GetPermssionDetailResponse {
        permission_id: permission.permission_id,
        object: permission.object,
        action: permission.action,
        name: permission.name,
        description: permission.description,
    };
    Ok(ApiOk(Some(permission)))
}

async fn get_permission_by_id(id: i64) -> Result<entity::permission::Model, ApiError> {
    let db = utils::db::conn();
    let permission = entity::permission::Entity::find_by_id(id)
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find permission");
            utils::response::ApiError::DbError(None)
        })?;

    if let Some(permission) = permission {
        Ok(permission)
    } else {
        Err(utils::response::ApiError::RequestError(Some("Permission not found".to_string())))
    }
}