use salvo::prelude::*;
use salvo::oapi::ToParameters;
use serde::{Deserialize, Serialize};
use utils::response::{ApiOk, ApiResult};

use crate::modules::file::service;


#[derive(Deserialize, ToParameters)]
#[salvo(parameters(default_parameter_in = Query))]
pub struct GetFileUrlRequest {
    pub path: String,
}


#[derive(Serialize, ToSchema)]
pub struct GetFileUrlResponse {
    pub url: String,
}

/// Get file's url by file path
#[endpoint(
    tags("File"),
)]
pub async fn get_file_url(
    query: GetFileUrlRequest,
) -> ApiResult<GetFileUrlResponse> {
    let sign_url = service::get::get_file_download_url(&query.path).await?;
    let response = GetFileUrlResponse {
        url: sign_url,
    };
    Ok(ApiOk(Some(response)))
}