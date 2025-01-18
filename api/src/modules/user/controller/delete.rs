use salvo::prelude::*;
use salvo::oapi::extract::PathParam;
use utils::response::{ApiOk, ApiResult};

use crate::modules::user::service;

/// Delete a user
#[endpoint(
    tags("User"),
)]
pub async fn delete_user(
    user_id: PathParam<i64>,
) -> ApiResult<bool> {
    let _ = service::delete::delete_user_by_id(user_id.into_inner()).await?;

    Ok(ApiOk(Some(true)))
}
