use salvo::{oapi::extract::{JsonBody, PathParam}, prelude::*};
use sea_orm::{EntityTrait, Set, ActiveModelTrait};
use serde::Deserialize;
use utils::response::{ApiError, ApiOk, ApiResult};
use validator::Validate;



#[derive(Debug, Deserialize, Validate, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct UpdateMenuRequest {
    pub name: String,
    pub parent_id: Option<i64>,
    pub icon: Option<String>,
    pub menu_type: i32,
    pub sort: i32,
    pub path: Option<String>,
}

/// Update a menu
#[endpoint(
    tags("Menu"),
)]
pub async fn update_menu(
    menu_id: PathParam<i64>,
    body: JsonBody<UpdateMenuRequest>,
) -> ApiResult<bool> {
    let _ = update_menu_by_request(menu_id.into_inner(), body.into_inner()).await?;

    Ok(ApiOk(Some(true)))
}

async fn update_menu_by_request(menu_id: i64, req: UpdateMenuRequest) -> Result<bool, ApiError> {
    let db = utils::db::conn();
    let menu = entity::menu::Entity::find_by_id(menu_id)
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find menu");
            ApiError::DbError(None)
        })?;

    if menu.is_none() {
        return Err(ApiError::RequestError(Some("Menu not found".to_string())));
    }

    let mut menu: entity::menu::ActiveModel = menu.unwrap().into();
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