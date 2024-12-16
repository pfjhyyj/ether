use sea_orm_migration::prelude::*;

use crate::{m20241019_115953_init_menu::Menu, m20241020_021506_init_role::Role};

#[derive(DeriveMigrationName)]
pub struct Migration;

#[async_trait::async_trait]
impl MigrationTrait for Migration {
    async fn up(&self, manager: &SchemaManager) -> Result<(), DbErr> {
        manager
            .create_table(
                Table::create()
                    .table(RoleMenu::Table)
                    .if_not_exists()
                    .col(
                        ColumnDef::new(RoleMenu::Id)
                            .big_integer()
                            .not_null()
                            .auto_increment()
                            .primary_key(),
                    )
                    .col(
                        ColumnDef::new(RoleMenu::RoleId)
                            .big_integer()
                            .not_null(),
                    )
                    .col(
                        ColumnDef::new(RoleMenu::MenuId)
                            .big_integer()
                            .not_null(),
                    )
                    .foreign_key(
                        ForeignKey::create()
                            .name("fk-role_menu-role_id")
                            .from(RoleMenu::Table, RoleMenu::RoleId)
                            .to(Role::Table, Role::RoleId)
                            .on_delete(ForeignKeyAction::Cascade)
                            .on_update(ForeignKeyAction::Cascade)
                    )
                    .foreign_key(
                        ForeignKey::create()
                            .name("fk-role_menu-menu_id")
                            .from(RoleMenu::Table, RoleMenu::MenuId)
                            .to(Menu::Table, Menu::MenuId)
                            .on_delete(ForeignKeyAction::Cascade)
                            .on_update(ForeignKeyAction::Cascade)
                    )
                    .to_owned(),
            )
            .await
    }

    async fn down(&self, manager: &SchemaManager) -> Result<(), DbErr> {
        manager
            .drop_table(Table::drop().table(RoleMenu::Table).to_owned())
            .await
    }
}

#[derive(DeriveIden)]
enum RoleMenu {
    Table,
    Id,
    RoleId,
    MenuId,
}
