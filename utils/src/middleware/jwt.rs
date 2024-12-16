use salvo::{async_trait, http::header::AUTHORIZATION, writing::Json, Depot, FlowCtrl, Handler, Request, Response};

use crate::{identity::Identity, response::ApiError};
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
            req.extensions_mut().insert(token_data);
        } else {
            resp.render(Json(ApiError::AuthError(Some("Invalid token".to_string())).to_response()));
            ctrl.skip_rest();
            return;
        }
    }
}