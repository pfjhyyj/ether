use salvo::Router;

pub mod controller;
pub mod service;

pub fn get_router() -> Router {
    Router::new()
        .path("/permissions")
        .push(Router::new().post(controller::create::create_permission).get(controller::list::page_permission))
        .push(Router::with_path("/{permission_id}").get(controller::get::get_permission).put(controller::update::update_permission).delete(controller::delete::delete_permission))
}