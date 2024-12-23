use salvo::{oapi::extract::PathParam, prelude::*};
use sea_orm::{EntityTrait, ModelTrait};
use utils::response::{ApiError, ApiOk, ApiResult};

/// Delete a role
#[endpoint(
    tags("Role"),
)]
pub async fn delete_role(
    role_id: PathParam<i64>,
) -> ApiResult<bool> {
    let _ = delete_role_by_id(role_id.into_inner()).await?;

    Ok(ApiOk(Some(true)))
}

async fn delete_role_by_id(role_id: i64) -> Result<bool, ApiError> {
    let db = utils::db::conn();
    let role = entity::role::Entity::find_by_id(role_id)
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find role");
            ApiError::DbError(None)
        })?;
    
    if let Some(role) = role {
        role.delete(db)
            .await
            .map_err(|e| {
                tracing::error!(error = ?e, "Failed to delete role");
                ApiError::DbError(None)
            })?;
    }
    Ok(true)
}