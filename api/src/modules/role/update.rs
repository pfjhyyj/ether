use salvo::{oapi::extract::{JsonBody, PathParam}, prelude::*};
use sea_orm::{EntityTrait, Set, ActiveModelTrait};
use domain::entity::role;
use serde::Deserialize;
use utils::response::{ApiError, ApiOk, ApiResult};
use validator::Validate;

#[derive(Debug, Deserialize, Validate, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct UpdateRoleRequest {
    pub code: String,
    pub reference_type: Option<String>,
    pub reference_id: Option<i64>,
    pub name: String,
    pub description: Option<String>,
}

/// Update a role
#[endpoint(
    tags("Role"),
)]
pub async fn update_role(
    role_id: PathParam<i64>,
    body: JsonBody<UpdateRoleRequest>,
) -> ApiResult<bool> {
    let _ = update_role_by_request(role_id.into_inner(), body.into_inner()).await?;

    Ok(ApiOk(Some(true)))
}

async fn update_role_by_request(role_id: i64, req: UpdateRoleRequest) -> Result<bool, ApiError> {
    let db = utils::db::conn();
    let role = role::Entity::find_by_id(role_id)
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find role");
            ApiError::DbError(None)
        })?;

    if role.is_none() {
        return Err(ApiError::RequestError(Some("Role not found".to_string())));
    }

    let mut role: role::ActiveModel = role.unwrap().into();
    role.code = Set(req.code);
    role.reference_type = Set(req.reference_type);
    role.reference_id = Set(req.reference_id);
    role.name = Set(req.name);
    role.description = Set(req.description);

    role.save(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to update role");
            ApiError::DbError(None)
        })?;

    Ok(true)
}