use sea_orm::{EntityTrait, Set, ActiveModelTrait};
use domain::entity::menu;
use utils::response::ApiError;

use crate::modules::menu::controller::update::UpdateMenuRequest;


pub async fn update_menu_by_request(menu_id: i64, req: UpdateMenuRequest) -> Result<bool, ApiError> {
    let db = utils::db::conn();
    let menu = menu::Entity::find_by_id(menu_id)
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find menu");
            ApiError::DbError(None)
        })?;

    if menu.is_none() {
        return Err(ApiError::RequestError(Some("Menu not found".to_string())));
    }

    let mut menu: menu::ActiveModel = menu.unwrap().into();
    menu.name = Set(req.name);
    menu.parent_id = Set(req.parent_id);
    menu.icon = Set(req.icon);
    menu.menu_type = Set(req.menu_type);
    menu.sort = Set(req.sort);
    menu.path = Set(req.path);


    menu.save(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to update menu");
            ApiError::DbError(None)
        })?;

    Ok(true)
}