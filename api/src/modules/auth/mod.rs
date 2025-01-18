use salvo::Router;

pub mod controller;
pub mod service;

pub fn get_open_router() -> Router {
    Router::new()
        .path("/auth")
        .push(Router::with_path("/login").post(controller::login::login_by_username))
        .push(Router::with_path("/register").post(controller::register::register_by_username))
}

pub fn get_router() -> Router {
    Router::new()
        .path("/auth")
        .push(Router::with_path("/logout").post(controller::logout::logout))
        .push(Router::with_path("/current").get(controller::current::get_current_user))
        .push(Router::with_path("/current").put(controller::current::update_current_user))
        .push(Router::with_path("/current/password").put(controller::current_password::update_password))
}
