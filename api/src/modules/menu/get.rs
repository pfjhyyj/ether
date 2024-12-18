use salvo::{oapi::extract::PathParam, prelude::*};
use sea_orm::EntityTrait;
use serde::Serialize;
use utils::response::{ApiError, ApiOk, ApiResult};

#[derive(Debug, Serialize, ToSchema)]
pub struct GetMenuDetailResponse {
    pub menu_id: i64,
    pub name: String,
    pub parent_id: Option<i64>,
    pub icon: Option<String>,
    pub menu_type: i32,
    pub sort: i32,
    pub path: Option<String>,
}

#[endpoint(
    tags("Menu"),
)]
pub async fn get_menu(
    menu_id: PathParam<i64>,
) -> ApiResult<GetMenuDetailResponse> {
    let menu = get_menu_by_id(menu_id.into_inner()).await?;
    let menu = GetMenuDetailResponse {
        menu_id: menu.menu_id,
        name: menu.name,
        parent_id: menu.parent_id,
        icon: menu.icon,
        menu_type: menu.menu_type,
        sort: menu.sort,
        path: menu.path,
    };
    Ok(ApiOk(Some(menu)))
}

async fn get_menu_by_id(id: i64) -> Result<entity::menu::Model, ApiError> {
    let db = utils::db::conn();
    let menu = entity::menu::Entity::find_by_id(id)
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