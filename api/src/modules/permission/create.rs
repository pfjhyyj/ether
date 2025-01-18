use salvo::{oapi::extract::JsonBody, prelude::*};
use sea_orm::{Set, ActiveModelTrait};
use domain::entity::permission;
use serde::Deserialize;
use utils::response::{ApiError, ApiOk, ApiResult};
use validator::Validate;


#[derive(Debug, Deserialize, Validate, ToSchema)]
pub struct CreatePermissionRequest {
    pub object: String,
    pub action: String,
    pub name: Option<String>,
    pub description: Option<String>,
}

/// Create a new permission
#[endpoint(
    tags("Permission"),
)]
pub async fn create_permission(
    body: JsonBody<CreatePermissionRequest>,
) -> ApiResult<i64> {
    let new_permission = create_permission_by_request(body.into_inner()).await?;

    Ok(ApiOk(Some(new_permission.permission_id)))
}

async fn create_permission_by_request(req: CreatePermissionRequest) -> Result<permission::Model, ApiError> {
    let db = utils::db::conn();

    let new_permission = permission::ActiveModel {
        object: Set(req.object),
        action: Set(req.action),
        name: Set(req.name),
        description: Set(req.description),
        ..Default::default()
    }.insert(db);

    let new_permission = new_permission.await.map_err(|e| {
        tracing::error!(error = ?e, "Failed to insert new permission");
        ApiError::DbError(None)
    })?;

    Ok(new_permission)
}