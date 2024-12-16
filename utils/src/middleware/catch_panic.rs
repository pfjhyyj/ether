use std::panic::AssertUnwindSafe;

use futures::FutureExt;
use salvo::{async_trait, writing::Json, Depot, FlowCtrl, Handler, Request, Response};

use crate::response::ApiError;

pub struct CatchPanic;

impl CatchPanic {
    #[inline]
    pub fn new() -> Self {
        CatchPanic {}
    }
}

#[async_trait]
impl Handler for CatchPanic {
    async fn handle(
        &self,
        req: &mut Request,
        depot: &mut Depot,
        resp: &mut Response,
        ctrl: &mut FlowCtrl,
    ) {
        if AssertUnwindSafe(ctrl.call_next(req, depot, resp))
            .catch_unwind()
            .await.is_err()
        {
            resp.render(Json(ApiError::UnknownError(None).to_response()));
        }
    }
}