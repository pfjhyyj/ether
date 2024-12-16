use std::collections::HashMap;

use http::header::CONTENT_TYPE;
use http_body_util::BodyExt;
use hyper::HeaderMap;
use salvo::{
    async_trait, http::ReqBody, writing::Json, Depot, FlowCtrl, Handler, Request, Response,
};

use crate::{response::ApiError, xtime};

pub struct LogMiddleware;

impl LogMiddleware {
    #[inline]
    pub fn new() -> Self {
        LogMiddleware {}
    }
}

#[async_trait]
impl Handler for LogMiddleware {
    async fn handle(
        &self,
        req: &mut Request,
        depot: &mut Depot,
        resp: &mut Response,
        ctrl: &mut FlowCtrl,
    ) {
        let enter_time = xtime::now(None);
        let req_method = req.method().to_string();
        let req_uri = req.uri().to_string();
        // get body
        let (body, code) = drain_body(req).await;
        if let Some(v) = code {
            resp.render(Json(v.to_response()));
            ctrl.skip_rest();
            return;
        }
        ctrl.call_next(req, depot, resp).await;
        // calculate process time
        let duration = (xtime::now(None) - enter_time).to_string();
        tracing::info!(
            method = req_method,
            uri = req_uri,
            body = body,
            duration = duration,
            "Request info"
        );
    }
}

#[allow(dead_code)]
fn header_to_string(h: &HeaderMap) -> String {
    let mut map: HashMap<String, Vec<String>> = HashMap::new();
    for k in h.keys() {
        let mut vals: Vec<String> = Vec::new();
        for v in h.get_all(k) {
            if let Ok(s) = v.to_str() {
                vals.push(s.to_string())
            }
        }
        map.insert(k.to_string(), vals);
    }
    serde_json::to_string(&map).unwrap_or_else(|_| String::from("<none>"))
}

async fn drain_body(req: &mut Request) -> (Option<String>, Option<ApiError>) {
    let ok = match req.header::<String>(CONTENT_TYPE) {
        Some(v) => {
            v.starts_with("application/json") || v.starts_with("application/x-www-form-urlencoded")
        }
        None => false,
    };
    if !ok {
        return (None, None);
    }
    // get body
    let body = req.take_body();
    let bytes = match body.collect().await {
        Ok(v) => v.to_bytes(),
        Err(e) => {
            tracing::error!(err = ?e, "body.collect");
            return (None, Some(ApiError::UnknownError(None)));
        }
    };
    let body_str = std::str::from_utf8(&bytes).map(|s| s.to_string()).ok();
    // reset body
    req.replace_body(ReqBody::Once(bytes));
    (body_str, None)
}