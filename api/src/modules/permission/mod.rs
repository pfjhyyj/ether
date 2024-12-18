use salvo::Router;

pub mod create;
pub mod update;
pub mod delete;
pub mod get;
pub mod list;

pub fn get_router() -> Router {
    Router::new()
        .path("/permissions")
        .push(Router::with_path("").post(create::create_permission).get(list::page_permission))
        .push(Router::with_path("/<permission_id>").get(get::get_permission).put(update::update_permission).delete(delete::delete_permission))
}