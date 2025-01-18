use salvo::{oapi::extract::JsonBody, prelude::*};
use sea_orm::{Set, ActiveModelTrait};
use domain::entity::role;
use serde::Deserialize;
use utils::response::{ApiError, ApiOk, ApiResult};
use validator::Validate;



#[derive(Debug, Deserialize, Validate, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct CreateRoleRequest {
    pub code: String,
    pub reference_type: Option<String>,
    pub reference_id: Option<i64>,
    pub name: String,
    pub description: Option<String>,
}

/// Create a new role
#[endpoint(
    tags("Role"),
)]
pub async fn create_role(
    body: JsonBody<CreateRoleRequest>
) -> ApiResult<i64> {
    let new_role = create_role_by_request(body.into_inner()).await?;

    Ok(ApiOk(Some(new_role.role_id)))
}

async fn create_role_by_request(req: CreateRoleRequest) -> Result<role::Model, ApiError> {
    let db = utils::db::conn();

    let new_role = role::ActiveModel {
        code: Set(req.code),
        reference_type: Set(req.reference_type),
        reference_id: Set(req.reference_id),
        name: Set(req.name),
        description: Set(req.description),
        ..Default::default()
    }.insert(db);

    let new_role = new_role.await.map_err(|e| {
        tracing::error!(error = ?e, "Failed to insert new role");
        ApiError::DbError(None)
    })?;

    Ok(new_role)
}