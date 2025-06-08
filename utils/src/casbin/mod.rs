use std::sync::OnceLock;

use casbin::{CoreApi, DefaultModel, Enforcer};

pub mod adapter;
pub mod model;

static CASBIN_CONNECTION: OnceLock<Enforcer> = OnceLock::new();

pub async fn init_casbin() {
    let model_path = "config/casbin_model.conf";

    let model = DefaultModel::from_file(model_path)
        .await
        .unwrap_or_else(|e| panic!("Failed to load Casbin model: {}", e));

    let adapter = adapter::MemoryAdapter::new(vec![]);

    let enforcer = Enforcer::new(model, adapter)
        .await
        .unwrap_or_else(|e| panic!("Failed to create Casbin enforcer: {}", e));

    let _ = CASBIN_CONNECTION.set(enforcer);
}

pub fn get_casbin() -> &'static Enforcer {
    CASBIN_CONNECTION
        .get()
        .unwrap_or_else(|| panic!("Casbin enforcer is not initialized"))
}