use salvo::prelude::*;
use domain::entity::user;
use salvo::oapi::extract::PathParam;
use sea_orm::{EntityTrait, ModelTrait};
use utils::response::{ApiError, ApiOk, ApiResult};

/// Delete a user
#[endpoint(
    tags("User"),
)]
pub async fn delete_user(
    user_id: PathParam<i64>,
) -> ApiResult<bool> {
    let _ = delete_user_by_id(user_id.into_inner()).await?;

    Ok(ApiOk(Some(true)))
}

async fn delete_user_by_id(user_id: i64) -> Result<bool, ApiError> {
    let db = utils::db::conn();
    let user = user::Entity::find_by_id(user_id)
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find user");
            ApiError::DbError(None)
        })?;

    if let Some(user) = user {
        user.delete(db)
            .await
            .map_err(|e| {
                tracing::error!(error = ?e, "Failed to delete user");
                ApiError::DbError(None)
            })?;
    }

    Ok(true)
}