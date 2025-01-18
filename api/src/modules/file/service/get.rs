use utils::response::ApiError;



pub async fn get_file_download_url(key: &str) -> Result<String, ApiError> {
    let signed_url = utils::s3::generate_file_download_params(key)
        .await?;

    Ok(signed_url)
}