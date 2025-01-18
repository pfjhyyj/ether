use salvo::Router;

pub mod controller;
pub mod service;

pub fn get_router() -> Router {
    Router::new()
        .path("/menus")
        .push(Router::new().post(controller::create::create_menu).get(controller::list::list_menu))
        .push(Router::with_path("/{menu_id}").get(controller::get::get_menu).put(controller::update::update_menu).delete(controller::delete::delete_menu))
}