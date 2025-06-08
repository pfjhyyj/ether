use sea_orm_migration::prelude::*;

use crate::{m20241020_021506_init_role::Role, m20241020_025112_init_user_role::UserRole};

#[derive(DeriveMigrationName)]
pub struct Migration;

#[async_trait::async_trait]
impl MigrationTrait for Migration {
    async fn up(&self, manager: &SchemaManager) -> Result<(), DbErr> {
        let insert = Query::insert()
            .into_table(Role::Table)
            .columns([
                Role::RoleId,
                Role::Code,
                Role::Name,
                Role::Description,
            ])
            .values_panic(vec![
                1.into(),
                "admin".into(),
                "Admin".into(),
                "Administrator".into(),
            ])
            .values_panic(vec![
                2.into(),
                "user".into(),
                "User".into(),
                "Normal User".into(),
            ])
            .to_owned();
        manager.exec_stmt(insert).await?;

        let insert = Query::insert()
            .into_table(UserRole::Table)
            .columns([UserRole::UserId, UserRole::RoleId])
            .values_panic(vec![1.into(), 1.into()])
            .to_owned();
        manager.exec_stmt(insert).await?;
        Ok(())
    }

    async fn down(&self, manager: &SchemaManager) -> Result<(), DbErr> {
        let delete = Query::delete()
            .from_table(Role::Table)
            .to_owned();
        manager.exec_stmt(delete).await?;

        let delete = Query::delete()
            .from_table(UserRole::Table)
            .to_owned();
        manager.exec_stmt(delete).await?;
        Ok(())
    }
}