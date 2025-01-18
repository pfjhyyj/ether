use salvo::{oapi::extract::PathParam, prelude::*};
use serde::Serialize;
use utils::response::{ApiOk, ApiResult};

use crate::modules::permission::service;

#[derive(Debug, Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct GetPermssionDetailResponse {
    pub permission_id: i64,
    pub object: String,
    pub action: String,
    pub name: Option<String>,
    pub description: Option<String>,
}

/// Get a permission
#[endpoint(
    tags("Permission"),
)]
pub async fn get_permission(
    permission_id: PathParam<i64>,
) -> ApiResult<GetPermssionDetailResponse> {
    let permission = service::get::get_permission_by_id(permission_id.into_inner()).await?;
    let permission = GetPermssionDetailResponse {
        permission_id: permission.permission_id,
        object: permission.object,
        action: permission.action,
        name: permission.name,
        description: permission.description,
    };
    Ok(ApiOk(Some(permission)))
}
