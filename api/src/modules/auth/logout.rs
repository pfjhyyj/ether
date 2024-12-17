use salvo::prelude::*;
use redis::Commands;
use salvo::Request;
use utils::{identity::Identity, response::{ApiError, ApiOk, ApiResult}};

#[endpoint(
    tags("Auth"),
)]
pub async fn logout(
    req: &mut Request
) -> ApiResult<()> {
    let id = req.extensions().get::<Identity>().unwrap();
    clear_token_cache(id.sub)?;
    Ok(ApiOk(None))
}

fn clear_token_cache(user_id: i64) -> Result<(), ApiError> {
    let mut conn = match utils::redis::redis_pool().get() {
        Ok(c) => c,
        Err(e) => {
            tracing::error!(error = ?e, "Failed to get redis connection");
            return Err(ApiError::UnknownError(None));
        }
    };

    let key = format!("token:{}", user_id);
    let _: () = conn.del(key).map_err(|e| {
        tracing::error!(error = ?e, "Failed to clear token cache");
        ApiError::UnknownError(None)
    })?;

    Ok(())
}
