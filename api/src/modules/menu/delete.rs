use salvo::{oapi::extract::PathParam, prelude::*};
use domain::entity::menu;
use sea_orm::{EntityTrait, ModelTrait};
use utils::response::{ApiError, ApiOk, ApiResult};

/// Delete a menu
#[endpoint(
    tags("Menu"),
)]
pub async fn delete_menu(
    menu_id: PathParam<i64>,
) -> ApiResult<bool> {
    let _ = delete_menu_by_id(menu_id.into_inner()).await?;

    Ok(ApiOk(Some(true)))
}


async fn delete_menu_by_id(menu_id: i64) -> Result<bool, ApiError> {
    let db = utils::db::conn();
    let menu = menu::Entity::find_by_id(menu_id)
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find menu");
            ApiError::DbError(None)
        })?;
    
    if let Some(menu) = menu {
        menu.delete(db)
            .await
            .map_err(|e| {
                tracing::error!(error = ?e, "Failed to delete menu");
                ApiError::DbError(None)
            })?;
    }
    Ok(true)
}