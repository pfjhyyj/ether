use salvo::{oapi::extract::PathParam, prelude::*};
use utils::response::{ApiOk, ApiResult};

use crate::modules::role::service;

/// Delete a role
#[endpoint(
    tags("Role"),
)]
pub async fn delete_role(
    role_id: PathParam<i64>,
) -> ApiResult<bool> {
    service::delete::delete_role_by_id(role_id.into_inner()).await?;

    Ok(ApiOk(Some(true)))
}
