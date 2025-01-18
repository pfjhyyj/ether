use salvo::Router;

pub mod controller;
pub mod service;

pub fn get_router() -> Router {
    Router::new()
        .path("/users")
        .push(Router::new().get(controller::list::page_user).post(controller::create::create_user))
        .push(Router::with_path("/{user_id}").get(controller::get::get_user).put(controller::update::update_user).delete(controller::delete::delete_user))
}
