use salvo::{oapi::extract::{JsonBody, PathParam}, prelude::*};
use serde::Deserialize;
use utils::response::{ApiError, ApiOk, ApiResult};
use sea_orm::{ColumnTrait, EntityTrait, QueryFilter, Set};
use validator::Validate;
use domain::entity::{role, permission, role_permission};


#[derive(Debug, Deserialize, Validate, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct AddRolePermissionsRequest {
    pub permission_ids: Vec<i64>,
}

/// Add permissions to a role
#[endpoint(
    tags("Role"),
)]
pub async fn add_role_permissions(
    role_id: PathParam<i64>,
    body: JsonBody<AddRolePermissionsRequest>
) -> ApiResult<bool> {
    let _ = add_role_permissions_by_request(role_id.into_inner(), body.into_inner()).await?;
    
    Ok(ApiOk(Some(true)))
}

async fn add_role_permissions_by_request(role_id: i64, req: AddRolePermissionsRequest) -> Result<bool, ApiError> {
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

    // check if permissions exist
    let permissions = permission::Entity::find()
        .filter(permission::Column::PermissionId.is_in(req.permission_ids.clone()))
        .all(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find permissions");
            ApiError::DbError(None)
        })?;
    
    if permissions.len() != req.permission_ids.len() {
        return Err(ApiError::RequestError(Some("Permission not found".to_string())));
    }

    let role_permissions = role_permission::Entity::find()
        .filter(role_permission::Column::RoleId.eq(role_id))
        .all(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find role permissions");
            ApiError::DbError(None)
        })?;
    
    // check if permissions already exist
    let role_permission_ids: Vec<i64> = role_permissions.iter().map(|rp| rp.permission_id).collect();
    // get new permissions
    let new_permissions: Vec<i64> = req.permission_ids.iter().filter(|p| !role_permission_ids.contains(p)).map(|p| *p).collect();
    
    if new_permissions.is_empty() {
        return Err(ApiError::RequestError(Some("Permissions already exist".to_string())));
    }

    // insert new permissions
    let new_role_permissions = new_permissions.iter().map(|p| {
        role_permission::ActiveModel {
            role_id: Set(role_id),
            permission_id: Set(*p),
            ..Default::default()
        }
    }).collect::<Vec<_>>();

    let _ = role_permission::Entity::insert_many(new_role_permissions)
        .exec(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to insert new role permissions");
            ApiError::DbError(None)
        })?;

    Ok(true)
}