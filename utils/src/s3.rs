use std::borrow::Cow;

use s3::{creds::Credentials, Bucket, PostPolicy, PostPolicyField, Region};

use crate::{config, response::ApiError};

pub struct FileUploadParams {
    pub x_amz_credential: String,
    pub x_amz_date: String,
    pub x_amz_algorithm: String,
    pub x_amz_signature: String,
    pub policy: String,
    pub key: String,
    pub bucket: String,
    pub host: String,
}

pub async fn generate_file_upload_params(key: &str) -> Result<FileUploadParams, ApiError> {
    let access_key = config::global().get_string("s3.access_key_id").unwrap_or_else(|e| panic!("Failed to get s3.access_key_id: {}", e));
    let secret_key = config::global().get_string("s3.access_key_serect").unwrap_or_else(|e| panic!("Failed to get s3.secret: {}", e));
    let region_name = config::global().get_string("s3.region").unwrap_or_else(|e| panic!("Failed to get s3.region: {}", e));
    let endpoint = config::global().get_string("s3.endpoint").unwrap_or_else(|e| panic!("Failed to get s3.endpoint: {}", e));
    let bucket_name = config::global().get_string("s3.bucket").unwrap_or_else(|e| panic!("Failed to get s3.bucket: {}", e));


    let credentials = Credentials::new(Some(&access_key), Some(&secret_key), None, None, Some("Static")).unwrap();
    let region = Region::Custom { region: region_name, endpoint: endpoint };
    let bucket = Bucket::new(&bucket_name, region, credentials).map_err(|err| ApiError::UnknownError(Some(err.to_string())))?;

    let post_policy = PostPolicy::new(60 * 10).condition(
        PostPolicyField::Key,
        s3::PostPolicyValue::Exact(Cow::from(key))
    ).map_err(|err| ApiError::UnknownError(Some(err.to_string())))?;
    
    let presigned_post = bucket.presign_post(post_policy).await.map_err(|err| ApiError::UnknownError(Some(err.to_string())))?;
    println!("Presigned url: {}, fields: {:?}", presigned_post.url, presigned_post.fields);
    Ok(FileUploadParams {
        x_amz_credential: presigned_post.fields.get("x-amz-credential").unwrap().to_string(),
        x_amz_date: presigned_post.fields.get("x-amz-date").unwrap().to_string(),
        x_amz_algorithm: presigned_post.fields.get("x-amz-algorithm").unwrap().to_string(),
        x_amz_signature: presigned_post.fields.get("x-amz-signature").unwrap().to_string(),
        policy: presigned_post.fields.get("Policy").unwrap().to_string(),
        key: key.to_string(),
        bucket: bucket_name.to_string(),
        host: presigned_post.url
    })
}


pub async fn generate_file_download_params(key: &str) -> Result<String, ApiError> {
    let access_key = config::global().get_string("s3.access_key_id").unwrap_or_else(|e| panic!("Failed to get s3.access_key_id: {}", e));
    let secret_key = config::global().get_string("s3.access_key_serect").unwrap_or_else(|e| panic!("Failed to get s3.secret: {}", e));
    let region_name = config::global().get_string("s3.region").unwrap_or_else(|e| panic!("Failed to get s3.region: {}", e));
    let endpoint = config::global().get_string("s3.endpoint").unwrap_or_else(|e| panic!("Failed to get s3.endpoint: {}", e));
    let bucket_name = config::global().get_string("s3.bucket").unwrap_or_else(|e| panic!("Failed to get s3.bucket: {}", e));


    let credentials = Credentials::new(Some(&access_key), Some(&secret_key), None, None, Some("Static")).unwrap();
    let region = Region::Custom { region: region_name, endpoint: endpoint };
    let bucket = Bucket::new(&bucket_name, region, credentials).map_err(|err| ApiError::UnknownError(Some(err.to_string())))?;

    let url: String = bucket.presign_get(key, 60 * 60, None).await.map_err(|err| ApiError::UnknownError(Some(err.to_string())))?;
    Ok(url)
}