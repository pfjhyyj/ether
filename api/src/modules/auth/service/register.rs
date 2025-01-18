use domain::entity::user;
use sea_orm::{ActiveModelTrait, ColumnTrait, EntityTrait, QueryFilter, Set};
use utils::response::ApiError;

use crate::modules::auth::controller::register::RegisterByUsernameRequest;

pub async fn create_user_by_register_request(req: RegisterByUsernameRequest) -> Result<user::Model, ApiError> {
    let db = utils::db::conn();
    let user = user::Entity::find()
        .filter(user::Column::Username.eq(&req.username))
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to query user by username");
            ApiError::DbError(None)
        })?;

    if user.is_some() {
        return Err(ApiError::RequestError(Some("Username already exists".to_string())));
    }

    let password = utils::hash::bcrypt(&req.password);

    let new_user = user::ActiveModel {
        username: Set(req.username),
        password: Set(password),
        ..Default::default()
    };

    let new_user = new_user.insert(db).await.map_err(|e| {
        tracing::error!(error = ?e, "Failed to insert new user");
        ApiError::DbError(None)
    })?;

    Ok(new_user)
}
