use salvo::{oapi::extract::PathParam, prelude::*};
use serde::Serialize;
use utils::response::{ApiOk, ApiResult};

use crate::modules::menu::service;

#[derive(Debug, Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct GetMenuDetailResponse {
    pub menu_id: i64,
    pub name: String,
    pub parent_id: Option<i64>,
    pub icon: Option<String>,
    pub menu_type: i32,
    pub sort: i32,
    pub path: Option<String>,
}

/// Get a menu
#[endpoint(
    tags("Menu"),
)]
pub async fn get_menu(
    menu_id: PathParam<i64>,
) -> ApiResult<GetMenuDetailResponse> {
    let menu = service::get::get_menu_by_id(menu_id.into_inner()).await?;
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

