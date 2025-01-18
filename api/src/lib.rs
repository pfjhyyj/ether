use salvo::prelude::*;

mod modules;

#[tokio::main]
pub async fn start(config_file_path: &str) -> std::io::Result<()> {
    utils::config::init(&config_file_path);
    let _guard = utils::logger::init();
    utils::db::init_db().await;
    utils::redis::init_redis();

    let port = utils::config::global().get_int("app.port").unwrap_or(3456);
    let address = format!("0.0.0.0:{}", port);

    let mut router = Router::new().push(modules::get_router());

    let show_openapi = utils::config::global().get_bool("app.debug").unwrap_or(false);
    if show_openapi {
        let doc = OpenApi::new("Ether Api", "0.1.0").merge_router(&router);
        router = router
            .push(doc.into_router("/api-doc/openapi.json"))
            .push(SwaggerUi::new("/api-doc/openapi.json").into_router("swagger-ui"));
    }

    let acceptor = TcpListener::new(address).bind().await;
    Server::new(acceptor).serve(router).await;
    Ok(())
}


pub fn serve() {
    let profile = utils::env::get_env_with_default("ENV", "local".to_string());
    // the config file will be config.{}.toml, where {} is the value of the ENV environment variable
    let config_file_path = format!("config/config.{}.toml", profile);
    let result = start(&config_file_path);

    if let Some(err) = result.err() {
        tracing::error!("Failed to start server: {:?}", err);
    }
}