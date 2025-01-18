use salvo::{oapi::extract::PathParam, prelude::*};
use utils::response::{ApiOk, ApiResult};

use crate::modules::permission::service;

/// Delete a permission
#[endpoint(
    tags("Permission"),
)]
pub async fn delete_permission(
   permission_id: PathParam<i64>,
) -> ApiResult<bool> {
    let _ = service::delete::delete_permission_by_id(permission_id.into_inner()).await?;

    Ok(ApiOk(Some(true)))
}
