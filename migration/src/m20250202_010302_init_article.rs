use sea_orm_migration::prelude::*;

#[derive(DeriveMigrationName)]
pub struct Migration;

#[async_trait::async_trait]
impl MigrationTrait for Migration {
    async fn up(&self, manager: &SchemaManager) -> Result<(), DbErr> {
        manager
            .create_table(
                Table::create()
                    .table(Article::Table)
                    .if_not_exists()
                    .col(
                        ColumnDef::new(Article::ArticleId)
                            .big_integer()
                            .not_null()
                            .auto_increment()
                            .primary_key(),
                    )
                    .col(ColumnDef::new(Article::Title).string().not_null())
                    .col(ColumnDef::new(Article::Content).text().null())
                    .col(ColumnDef::new(Article::Category).string().null())
                    .col(ColumnDef::new(Article::Tag).string().null())
                    .col(
                        ColumnDef::new(Article::Status)
                            .integer()
                            .not_null()
                            .default(0),
                    )
                    .col(ColumnDef::new(Article::CreatedAt).timestamp().not_null().default(Expr::current_timestamp()))
                    .col(ColumnDef::new(Article::UpdatedAt).timestamp().not_null().default(Expr::current_timestamp()))
                    .col(ColumnDef::new(Article::DeletedAt).timestamp().null())
                    .to_owned(),
            )
            .await
    }

    async fn down(&self, manager: &SchemaManager) -> Result<(), DbErr> {
        manager
            .drop_table(Table::drop().table(Article::Table).to_owned())
            .await
    }
}

#[derive(DeriveIden)]
enum Article {
    Table,
    ArticleId,
    Title,
    Content,
    Category,
    Tag,
    Status,
    CreatedAt,
    UpdatedAt,
    DeletedAt,
}
