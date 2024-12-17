use entity::user;
use redis::Commands;
use salvo::prelude::*;
use salvo::oapi::{endpoint, extract::JsonBody, ToSchema};
use sea_orm::{ColumnTrait, EntityTrait, QueryFilter};
use serde::{Deserialize, Serialize};
use utils::response::{ApiError, ApiOk, ApiResult};
use utils::xtime;
use validator::Validate;

#[derive(Debug, Deserialize, ToSchema, Validate)]
pub struct LoginByUserNameRequest {
    #[validate(length(
        min = 6,
        max = 50,
        message = "Username must be between 6 and 50 characters"
    ))]
    pub username: String,
    #[validate(length(
        min = 6,
        max = 50,
        message = "Password must be between 6 and 50 characters"
    ))]
    pub password: String,
}

#[derive(Debug, Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct LoginByUserNameResponse {
    pub access_token: String,
    pub expire_time: i64,
}

#[endpoint(
    tags("Auth"),
)]
pub async fn login_by_username(
    req: JsonBody<LoginByUserNameRequest>,
) -> ApiResult<LoginByUserNameResponse> {
    let user = get_by_username(&req.username).await?;

    let is_valid = utils::hash::bcrypt_verify(&req.password, &user.password);
    if !is_valid {
        return Err(ApiError::RequestError(Some("Invalid username or password".to_string())));
    }

    //  7 days to expire
    let expire_time = xtime::now(None).unix_timestamp() + 60 * 60 * 24 * 7;
    let identity = utils::identity::Identity {
        sub: user.user_id,
        exp: expire_time,
    };
    let token = utils::jwt::generate_jwt_token(&identity).map_err(|e| {
        tracing::error!(error = ?e, "Failed to generate jwt token");
        ApiError::UnknownError(None)
    })?;

    set_token_cache(&token, user.user_id)?;

    let resp = LoginByUserNameResponse {
        access_token: token,
        expire_time: expire_time,
    };
    Ok(ApiOk(Some(resp)))
}

async fn get_by_username(username: &str) -> Result<user::Model, ApiError> {
    let db = utils::db::conn();
    let user = user::Entity::find()
        .filter(user::Column::Username.eq(username))
        .one(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to get user by username");
            ApiError::DbError(None)
        })?
        .ok_or(ApiError::RequestError(Some("Invalid username or password".to_string())))?;

    Ok(user)
}

fn set_token_cache(token: &str, user_id: i64) -> Result<(), ApiError> {
    let mut conn = match utils::redis::redis_pool().get() {
        Ok(c) => c,
        Err(e) => {
            tracing::error!(error = ?e, "Failed to get redis connection");
            return Err(ApiError::UnknownError(None));
        }
    };

    let key = format!("token:{}", user_id);
    let _: () = conn.set(key, token).map_err(|e| {
        tracing::error!(error = ?e, "Failed to set token cache");
        ApiError::UnknownError(None)
    })?;

    Ok(())
}
