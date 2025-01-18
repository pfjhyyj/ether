use salvo::Router;

mod create;
mod update;
mod delete;
mod get;
mod list;
mod add_permissions;
mod remove_permissions;
mod list_permissions;

pub fn get_router() -> Router {
    Router::new()
        .path("/roles")
        .push(Router::new().post(create::create_role).get(list::page_role))
        .push(Router::with_path("/{role_id}").get(get::get_role).put(update::update_role).delete(delete::delete_role))
        .push(Router::with_path("/{role_id}/permissions/add").post(add_permissions::add_role_permissions))
        .push(Router::with_path("/{role_id}/permissions/remove").post(remove_permissions::remove_role_permissions))
        .push(Router::with_path("/{role_id}/permissions").get(list_permissions::page_role_permissions))
}