use salvo::{oapi::extract::{JsonBody, PathParam}, prelude::*};
use serde::Deserialize;
use utils::response::{ApiOk, ApiResult};
use validator::Validate;

use crate::modules::menu::service;



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
    service::update::update_menu_by_request(menu_id.into_inner(), body.into_inner()).await?;

    Ok(ApiOk(Some(true)))
}
