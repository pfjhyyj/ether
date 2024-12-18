use salvo::{oapi::extract::{JsonBody, PathParam}, prelude::*};
use sea_orm::{EntityTrait, Set, ActiveModelTrait};
use serde::Deserialize;
use utils::response::{ApiError, ApiOk, ApiResult};
use validator::Validate;



#[derive(Debug, Deserialize, Validate, ToSchema)]
pub struct UpdatePermissionRequest {
    pub object: String,
    pub action: String,
    pub name: Option<String>,
    pub description: Option<String>,
}

#[endpoint(
    tags("Permission"),
)]
pub async fn update_permission(
    permission_id: PathParam<i64>,
    body: JsonBody<UpdatePermissionRequest>,
) -> ApiResult<bool> {
    let _ = update_permission_by_request(permission_id.into_inner(), body.into_inner()).await?;

    Ok(ApiOk(Some(true)))
}

async fn update_permission_by_request(permission_id: i64, req: UpdatePermissionRequest) -> Result<bool, ApiError> {
    let db = utils::db::conn();
    let permission = entity::permission::Entity::find_by_id(permission_id)
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find permission");
            ApiError::DbError(None)
        })?;

    if permission.is_none() {
        return Err(ApiError::RequestError(Some("Permission not found".to_string())));
    }

    let mut permission: entity::permission::ActiveModel = permission.unwrap().into();
    permission.object = Set(req.object);
    permission.action = Set(req.action);
    permission.name = Set(req.name);
    permission.description = Set(req.description);

    permission.save(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to update permission");
            ApiError::DbError(None)
        })?;

    Ok(true)
}