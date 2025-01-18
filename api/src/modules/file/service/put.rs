use utils::{response::ApiError, s3::FileUploadParams};
use chrono::Utc;
use uuid::Uuid;

pub async fn get_file_upload_url(ext: &str) -> Result<FileUploadParams, ApiError> {
    let date = Utc::now().format("%Y-%m-%d").to_string();
    let uuid = Uuid::new_v4().to_string();
    let key = format!("study-helper/management/uploads/{}/{}.{}", date, uuid, ext);

    let signed_url = utils::s3::generate_file_upload_params(&key)
        .await?;

    Ok(signed_url)
}