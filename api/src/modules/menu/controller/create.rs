use salvo::prelude::*;
use salvo::oapi::{endpoint, extract::JsonBody, ToSchema};
use serde::Deserialize;
use utils::response::{ApiOk, ApiResult};
use validator::Validate;

use crate::modules::menu::service;


#[derive(Debug, Deserialize, Validate, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct CreateMenuRequest {
    pub name: String,
    pub parent_id: Option<i64>,
    pub icon: Option<String>,
    pub menu_type: i32,
    pub sort: i32,
    pub path: Option<String>,
}

/// Create a new menu
#[endpoint(
    tags("Menu"),
)]
pub async fn create_menu(
    body: JsonBody<CreateMenuRequest>,
) -> ApiResult<i64> {
    let new_menu = service::create::create_menu_by_request(body.into_inner()).await?;

    Ok(ApiOk(Some(new_menu.menu_id)))
}
