use std::{sync::OnceLock, time::Duration};

pub use sea_orm::DatabaseConnection;
use sea_orm::{ConnectOptions, Database as SeaOrmDatabase};

use crate::config;

static DB_CONNECTION: OnceLock<DatabaseConnection> = OnceLock::new();

pub async fn init_db() {
    let cfg = config::global();
    let url = cfg.get_string("db.dsn").unwrap_or_else(|e| panic!("Failed to get db.dsn: {}", e));

    let min_conns = cfg
        .get_int("db.options.min_conns")
        .unwrap_or(10);
    let max_conns = cfg
        .get_int("options.max_conns")
        .unwrap_or(20);
    let conn_timeout = cfg
        .get_int("options.conn_timeout")
        .unwrap_or(10);
    let idle_timeout = cfg
        .get_int("options.idle_timeout")
        .unwrap_or(300);
    let max_lifetime = cfg
        .get_int("options.max_lifetime")
        .unwrap_or(600);

    let mut opt = ConnectOptions::new(url);

    opt.min_connections(min_conns as u32)
        .max_connections(max_conns as u32)
        .connect_timeout(Duration::from_secs(conn_timeout as u64))
        .idle_timeout(Duration::from_secs(idle_timeout as u64))
        .max_lifetime(Duration::from_secs(max_lifetime as u64))
        .sqlx_logging(cfg.get_bool("app.debug").unwrap_or_default());

    let conn = SeaOrmDatabase::connect(opt)
        .await
        .unwrap_or_else(|e| panic!("Failed to connect to the database: {}", e));

    let _ = conn
        .ping()
        .await
        .is_err_and(|e| panic!("Failed to connect to the database: {}", e));

    let _ = DB_CONNECTION.set(conn);
}

pub fn conn() -> &'static DatabaseConnection {
    DB_CONNECTION
        .get()
        .unwrap_or_else(|| panic!("Database connection is not initiated"))
}
