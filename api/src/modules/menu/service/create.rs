use domain::entity::menu;
use sea_orm::{Set, ActiveModelTrait};
use utils::response::ApiError;

use crate::modules::menu::controller::create::CreateMenuRequest;

pub async fn create_menu_by_request(req: CreateMenuRequest) -> Result<menu::Model, ApiError> {
    let db = utils::db::conn();

    let new_menu = menu::ActiveModel {
        name: Set(req.name),
        parent_id: Set(req.parent_id),
        icon: Set(req.icon),
        menu_type: Set(req.menu_type),
        sort: Set(req.sort),
        path: Set(req.path),
        ..Default::default()
    }.insert(db);

    let new_menu = new_menu.await.map_err(|e| {
        tracing::error!(error = ?e, "Failed to insert new menu");
        ApiError::DbError(None)
    })?;

    Ok(new_menu)
}