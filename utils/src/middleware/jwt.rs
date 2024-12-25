use salvo::{async_trait, http::header::AUTHORIZATION, writing::Json, Depot, FlowCtrl, Handler, Request, Response};
use ::redis::Commands;

use crate::{cache::login::get_user_login_token_key, identity::Identity, redis, response::ApiError};
pub struct JwtMiddleware;

impl JwtMiddleware {
    #[inline]
    pub fn new() -> Self {
        JwtMiddleware {}
    }
}


#[async_trait]
impl Handler for JwtMiddleware {
    async fn handle(
        &self,
        req: &mut Request,
        _depot: &mut Depot,
        resp: &mut Response,
        ctrl: &mut FlowCtrl,
    ) {
        let token = req.header::<String>(AUTHORIZATION);
        let token_data = match token {
            Some(t) => Identity::from_auth_token(t),
            None => Identity::empty(),
        };

        if token_data.is_valid() {
            let key = get_user_login_token_key(token_data.sub);
            let mut conn = match redis::redis_pool().get() {
                Ok(c) => c,
                Err(e) => {
                    tracing::error!(error = ?e, "Failed to get redis connection");
                    resp.render(Json(ApiError::UnknownError(None).to_response()));
                    ctrl.skip_rest();
                    return;
                }
            };
            let token_result: Result<String, ApiError> = conn
                .get(key)
                .map_err(|_| ApiError::AuthError(Some("Token expired".to_string())));
            match token_result {
                Ok(token) => {
                    if token.is_empty() {
                        resp.render(Json(ApiError::AuthError(Some("Token expired".to_string())).to_response()));
                        ctrl.skip_rest();
                        return;
                    }
                }
                Err(e) => {
                    resp.render(Json(e.to_response()));
                    ctrl.skip_rest();
                    return;
                }
            }

            req.extensions_mut().insert(token_data);
        } else {
            resp.render(Json(ApiError::AuthError(Some("Token expired".to_string())).to_response()));
            ctrl.skip_rest();
            return;
        }
    }
}