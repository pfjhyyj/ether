use salvo::prelude::*;
use serde::Serialize;
use serde_json::Value;
use utils::{identity::Identity, response::{ApiOk, ApiResult}};

use crate::modules::menu::service;

#[derive(Debug, Serialize, Clone, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct MenuResponse {
    pub menu_id: i64,
    pub parent_id: Option<i64>,
    pub name: String,
    pub menu_type: i32,
    pub icon: Option<String>,
    pub path: Option<String>,
    pub sort: i32,
    pub extra: Option<Value>,
    pub children: Vec<MenuResponse>,
}

#[derive(Debug, Serialize, ToSchema)]
pub struct ListMenuResponse {
    pub menus: Vec<MenuResponse>,
}

/// List all menus
#[endpoint(
    tags("Menu"),
)]
pub async fn list_menu(
    req: &mut Request
) -> ApiResult<ListMenuResponse> {
    let _id = req.extensions().get::<Identity>().unwrap();
    let menus = service::list::get_menu_forest().await?;
    Ok(ApiOk(Some(menus)))
}
