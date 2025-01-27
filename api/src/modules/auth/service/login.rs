use domain::entity::user;
use redis::Commands;
use sea_orm::{ColumnTrait, EntityTrait, QueryFilter};
use utils::cache::login::get_user_login_token_key;
use utils::response::ApiError;
use utils::xtime;

use crate::modules::auth::controller::login::{LoginByUserNameRequest, LoginByUserNameResponse};


pub async fn login_by_username(
    req: LoginByUserNameRequest,
) -> Result<LoginByUserNameResponse, ApiError> {
    let user = get_by_username(&req.username).await?;

    let is_valid = utils::hash::bcrypt_verify(&req.password, &user.password);
    if !is_valid {
        return Err(ApiError::RequestError(Some("Invalid username or password".to_string())));
    }
    let valid_time_seconds: u64 = 60 * 60 * 24 * 7;
    //  7 days to expire
    let expire_time = xtime::now(None).unix_timestamp() + valid_time_seconds as i64;
    let identity = utils::identity::Identity {
        sub: user.user_id,
        exp: expire_time,
    };
    let token = utils::jwt::generate_jwt_token(&identity).map_err(|e| {
        tracing::error!(error = ?e, "Failed to generate jwt token");
        ApiError::UnknownError(None)
    })?;

    set_token_cache(&token, user.user_id, valid_time_seconds)?;

    let resp = LoginByUserNameResponse {
        access_token: token,
        expire_time: expire_time,
    };
    Ok(resp)
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

fn set_token_cache(token: &str, user_id: i64, valid_time_seconds: u64) -> Result<(), ApiError> {
    let mut conn = match utils::redis::redis_pool().get() {
        Ok(c) => c,
        Err(e) => {
            tracing::error!(error = ?e, "Failed to get redis connection");
            return Err(ApiError::UnknownError(None));
        }
    };

    let key = get_user_login_token_key(user_id);
    let _: () = conn.set_ex(key, token, valid_time_seconds).map_err(|e| {
        tracing::error!(error = ?e, "Failed to set token cache");
        ApiError::UnknownError(None)
    })?;

    Ok(())
}
