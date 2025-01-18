use std::collections::HashMap;

use domain::entity::menu;
use sea_orm::EntityTrait;
use utils::response::ApiError;

use crate::modules::menu::controller::list::{ListMenuResponse, MenuResponse};


pub async fn get_menu_forest() -> Result<ListMenuResponse, ApiError> {
    let menus = get_menu_list().await?;
    let menus = build_menu_forest(menus);
    let menus = ListMenuResponse { menus };
    Ok(menus)
}

async fn get_menu_list() -> Result<Vec<menu::Model>, ApiError> {
    let db = utils::db::conn();
    let menus = menu::Entity::find()
        .all(db)
        .await
        .map_err(|e| {
            tracing::error!(error = ?e, "Failed to find menus");
            utils::response::ApiError::DbError(None)
        })?;
    
    Ok(menus)
}

fn build_menu_forest(menus: Vec<menu::Model>) -> Vec<MenuResponse> {
    let mut menu_map: HashMap<Option<i64>, Vec<MenuResponse>> = HashMap::new();
    
    for menu in menus {
        let menu = MenuResponse {
            menu_id: menu.menu_id,
            parent_id: menu.parent_id,
            name: menu.name,
            menu_type: menu.menu_type,
            icon: menu.icon,
            path: menu.path,
            sort: menu.sort,
            extra: menu.extra,
            children: Vec::new(),
        };
        menu_map.entry(menu.parent_id).or_default().push(menu);
    }

    fn attach_children(parent_id: Option<i64>, menu_map: &mut HashMap<Option<i64>, Vec<MenuResponse>>) -> Vec<MenuResponse> {
        if let Some(children) = menu_map.remove(&parent_id) {
            children
                .into_iter()
                .map(|mut menu| {
                    menu.children = attach_children(Some(menu.menu_id), menu_map);
                    menu
                })
                .collect()
        } else {
            Vec::new()
        }
    }

    attach_children(None, &mut menu_map)
}