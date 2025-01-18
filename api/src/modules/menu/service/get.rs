use sea_orm::EntityTrait;
use domain::entity::menu;
use utils::response::ApiError;

pub async fn get_menu_by_id(id: i64) -> Result<menu::Model, ApiError> {
    let db = utils::db::conn();
    let menu = menu::Entity::find_by_id(id)
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find menu");
            utils::response::ApiError::DbError(None)
        })?;

    if let Some(menu) = menu {
        Ok(menu)
    } else {
        Err(utils::response::ApiError::RequestError(Some("Menu not found".to_string())))
    }
}