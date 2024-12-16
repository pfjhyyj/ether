use salvo::cors::Cors;
use salvo::oapi::extract::*;
use salvo::prelude::*;

pub mod auth;
// pub mod user;
// pub mod menu;
// pub mod permission;
// pub mod role;
#[endpoint]
async fn hello(name: QueryParam<String, false>) -> String {
    format!("Hello, {}!", name.as_deref().unwrap_or("World"))
}

pub fn get_router() -> Router {
    let open = Router::new()
        .hoop(utils::middleware::tracing::TracingMiddleware::new())
        // .get(hello)
        .push(auth::get_open_router());

    // let auth = Router::new()
    //     .hoop(utils::middleware::jwt::JwtMiddleware::new())
    //     .hoop(utils::middleware::tracing::TracingMiddleware::new())
    //     .push(auth::get_router());
        // .push(user::get_router())
        // .push(menu::get_router())
        // .push(permission::get_router())
        // .push(role::get_router());

    let cors = Cors::very_permissive()
        .expose_headers(vec![utils::middleware::tracing::TRACE_ID])
        .into_handler();

    Router::new()
        .hoop(cors)
        .hoop(utils::middleware::catch_panic::CatchPanic::new())
        .hoop(utils::middleware::log::LogMiddleware::new())
        .path("/api/v1")
        .push(open)
        // .push(auth)
}
