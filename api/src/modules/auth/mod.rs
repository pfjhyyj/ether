use salvo::Router;



pub mod login;
pub mod logout;
pub mod register;
pub mod current;
pub mod current_password;

pub fn get_open_router() -> Router {
    Router::new()
        .path("/auth")
        .push(Router::with_path("/login").post(login::login_by_username))
        // .push(Router::with_path("/register").post(register::register_by_username))
}

pub fn get_router() -> Router {
    Router::new()
        .path("/auth")
        .push(Router::with_path("/logout").post(logout::logout))
        .push(Router::with_path("/current").get(current::get_current_user))
        .push(Router::with_path("/current").put(current::update_current_user))
        .push(Router::with_path("/current/password").put(current_password::update_password))
}
