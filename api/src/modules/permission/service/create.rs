use sea_orm::{Set, ActiveModelTrait};
use domain::entity::permission;
use utils::response::ApiError;

use crate::modules::permission::controller::create::CreatePermissionRequest;


pub async fn create_permission_by_request(req: CreatePermissionRequest) -> Result<permission::Model, ApiError> {
    let db = utils::db::conn();

    let new_permission = permission::ActiveModel {
        object: Set(req.object),
        action: Set(req.action),
        name: Set(req.name),
        description: Set(req.description),
        ..Default::default()
    }.insert(db);

    let new_permission = new_permission.await.map_err(|e| {
        tracing::error!(error = ?e, "Failed to insert new permission");
        ApiError::DbError(None)
    })?;

    Ok(new_permission)
}