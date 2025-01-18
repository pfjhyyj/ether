use salvo::Router;

pub mod controller;
pub mod service;

pub fn get_router() -> Router {
    Router::new()
        .path("/files")
        .push(Router::new().get(controller::get::get_file_url).post(controller::put::get_file_upload_url))
}