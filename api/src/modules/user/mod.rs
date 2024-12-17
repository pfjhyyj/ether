use salvo::Router;

pub mod list;
pub mod create;
pub mod update;
pub mod delete;
pub mod get;

pub fn get_router() -> Router {
    Router::new()
        .path("/user")
        .push(Router::with_path("").get(list::page_user).post(create::create_user))
        .push(Router::with_path("/<user_id>").get(get::get_user).put(update::update_user).delete(delete::delete_user))
}
