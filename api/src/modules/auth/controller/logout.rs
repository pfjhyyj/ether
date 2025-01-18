use salvo::prelude::*;
use salvo::Request;
use utils::{identity::Identity, response::{ApiOk, ApiResult}};

use crate::modules::auth::service;

/// Logout
#[endpoint(
    tags("Auth"),
)]
pub async fn logout(
    req: &mut Request
) -> ApiResult<()> {
    let id = req.extensions().get::<Identity>().unwrap();
    service::logout::clear_token_cache(id.sub)?;
    Ok(ApiOk(None))
}
