use salvo::oapi::extract::JsonBody;
use salvo::prelude::*;
use salvo::oapi::ToSchema;
use serde::{Deserialize, Serialize};
use utils::response::{ApiOk, ApiResult};

use crate::modules::file::service;

#[derive(Deserialize, ToSchema)]
pub struct GetFileUploadUrlRequest {
    pub ext: String,
}

#[derive(Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct GetFileUploadUrlResponse {
    pub x_amz_credential: String,
    pub x_amz_date: String,
    pub x_amz_algorithm: String,
    pub x_amz_signature: String,
    pub policy: String,
    pub key: String,
    pub bucket: String,
    pub host: String,
}


/// Get file's upload url
#[endpoint(
    tags("File"),
)]
pub async fn get_file_upload_url(
    body: JsonBody<GetFileUploadUrlRequest>,
) -> ApiResult<GetFileUploadUrlResponse> {
    let sign_url = service::put::get_file_upload_url(&body.ext).await?;
    let response = GetFileUploadUrlResponse {
        x_amz_credential: sign_url.x_amz_credential,
        x_amz_date: sign_url.x_amz_date,
        x_amz_algorithm: sign_url.x_amz_algorithm,
        x_amz_signature: sign_url.x_amz_signature,
        policy: sign_url.policy,
        key: sign_url.key,
        bucket: sign_url.bucket,
        host: sign_url.host,
    };
    Ok(ApiOk(Some(response)))
}