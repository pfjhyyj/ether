use salvo::prelude::*;

mod modules;

#[tokio::main]
pub async fn start() -> std::io::Result<()> {
    utils::env::load_config();
    utils::logger::init();
    utils::db::init_db().await;
    utils::redis::init_redis();

    let port = utils::env::get_env::<String>("API_PORT");
    let address = format!("0.0.0.0:{}", port);

    let mut router = Router::new().push(modules::get_router());

    let show_openapi = utils::env::get_env::<bool>("SHOW_OPENAPI");
    if show_openapi {
        let doc = OpenApi::new("Ether api", "0.0.1").merge_router(&router);
        router = router
            .push(doc.into_router("/api-doc/openapi.json"))
            .push(SwaggerUi::new("/api-doc/openapi.json").into_router("swagger-ui"));
    }

    let acceptor = TcpListener::new(address).bind().await;
    Server::new(acceptor).serve(router).await;
    Ok(())
}


pub fn main() {
    let result = start();

    if let Some(err) = result.err() {
        tracing::error!("Failed to start server: {:?}", err);
    }
}