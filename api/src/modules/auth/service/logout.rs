use redis::Commands;
use utils::{cache::login::get_user_login_token_key, response::ApiError};

pub fn clear_token_cache(user_id: i64) -> Result<(), ApiError> {
    let mut conn = match utils::redis::redis_pool().get() {
        Ok(c) => c,
        Err(e) => {
            tracing::error!(error = ?e, "Failed to get redis connection");
            return Err(ApiError::UnknownError(None));
        }
    };

    let key = get_user_login_token_key(user_id);
    let _: () = conn.del(key).map_err(|e| {
        tracing::error!(error = ?e, "Failed to clear token cache");
        ApiError::UnknownError(None)
    })?;

    Ok(())
}
