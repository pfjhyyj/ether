use salvo::{oapi::extract::PathParam, prelude::*};
use utils::response::{ApiOk, ApiResult};

use crate::modules::menu::service;

/// Delete a menu
#[endpoint(
    tags("Menu"),
)]
pub async fn delete_menu(
    menu_id: PathParam<i64>,
) -> ApiResult<bool> {
    let _ = service::delete::delete_menu_by_id(menu_id.into_inner()).await?;

    Ok(ApiOk(Some(true)))
}

