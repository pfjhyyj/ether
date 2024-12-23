use salvo::{oapi::extract::PathParam, prelude::*};
use sea_orm::EntityTrait;
use serde::Serialize;
use utils::response::{ApiError, ApiOk, ApiResult};

#[derive(Debug, Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct GetRoleDetailResponse {
    pub role_id: i64,
    pub code: String,
    pub reference_type: Option<String>,
    pub reference_id: Option<i64>,
    pub name: String,
    pub description: Option<String>,
}

/// Get a role
#[endpoint(
    tags("Role"),
)]
pub async fn get_role(
    role_id: PathParam<i64>,
) -> ApiResult<GetRoleDetailResponse> {
    let role = get_role_by_id(role_id.into_inner()).await?;
    let role = GetRoleDetailResponse {
        role_id: role.role_id,
        code: role.code,
        reference_type: role.reference_type,
        reference_id: role.reference_id,
        name: role.name,
        description: role.description,
    };
    Ok(ApiOk(Some(role)))
}

async fn get_role_by_id(id: i64) -> Result<entity::role::Model, ApiError> {
    let db = utils::db::conn();
    let role = entity::role::Entity::find_by_id(id)
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find role");
            utils::response::ApiError::DbError(None)
        })?;

    if let Some(role) = role {
        Ok(role)
    } else {
        Err(utils::response::ApiError::RequestError(Some("Role not found".to_string())))
    }
}