use salvo::Router;

pub mod controller;
pub mod service;

pub fn get_router() -> Router {
    Router::new()
        .path("/roles")
        .push(Router::new().post(controller::create::create_role).get(controller::list::page_role))
        .push(Router::with_path("/{role_id}").get(controller::get::get_role).put(controller::update::update_role).delete(controller::delete::delete_role))
        .push(Router::with_path("/{role_id}/permissions/add").post(controller::add_permissions::add_role_permissions))
        .push(Router::with_path("/{role_id}/permissions/remove").post(controller::remove_permissions::remove_role_permissions))
        .push(Router::with_path("/{role_id}/permissions").get(controller::list_permissions::page_role_permissions))
}