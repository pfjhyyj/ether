use sea_orm::{Set, ActiveModelTrait};
use domain::entity::role;
use utils::response::ApiError;

use crate::modules::role::controller::create::CreateRoleRequest;

pub async fn create_role_by_request(req: CreateRoleRequest) -> Result<role::Model, ApiError> {
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