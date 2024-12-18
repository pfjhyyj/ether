use salvo::prelude::*;
use entity::menu;
use salvo::oapi::{endpoint, extract::JsonBody, ToSchema};
use sea_orm::{Set, ActiveModelTrait};
use serde::Deserialize;
use utils::response::{ApiError, ApiOk, ApiResult};
use validator::Validate;


#[derive(Debug, Deserialize, Validate, ToSchema)]
pub struct CreateMenuRequest {
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
pub async fn create_menu(
    body: JsonBody<CreateMenuRequest>,
) -> ApiResult<i64> {
    let new_menu = create_menu_by_request(body.into_inner()).await?;

    Ok(ApiOk(Some(new_menu.menu_id)))
}

async fn create_menu_by_request(req: CreateMenuRequest) -> Result<menu::Model, ApiError> {
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