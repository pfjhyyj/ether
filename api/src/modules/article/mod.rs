use salvo::Router;

pub mod controller;
pub mod service;

pub fn get_router() -> Router {
    Router::new()
        .path("/articles")
        .push(Router::new().post(controller::create::create_article).get(controller::list::page_article))
        .push(Router::with_path("/{article_id}").get(controller::get::get_article).put(controller::update::update_article).delete(controller::delete::delete_article))
}