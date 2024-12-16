use salvo::oapi::{BasicType, Components, Object, Operation, RefOr};
use salvo::prelude::*;
use salvo::{Depot, Request, Response, Writer};
use serde::Serialize;


pub enum ApiResponseCode {
  Ok = 0,
  UnknownError = 10001,
  DbError = 10002,

  RequestError = 20001,
  AuthError = 20002,
  AuthorizedError = 20003,
}

#[derive(Serialize)]
pub struct ApiResponse<T>{
  pub code: i32,
  pub message: Option<String>,
  #[serde(skip_serializing_if = "Option::is_none")]
  pub data: Option<T>,
}

impl<T> ToResponse for ApiResponse<T>
where
    T: ToSchema,
{
    fn to_response(components: &mut Components) -> RefOr<salvo::oapi::Response> {
        let schema = <T as ToSchema>::to_schema(components);
        let wrapper_schema = Object::new()
            .property("code", Object::new().description("response code").schema_type(BasicType::Integer))
            .property("message", Object::new().description("response message").schema_type(BasicType::String))
            .property("data", schema);
        salvo::oapi::Response::new("Response with json format data")
            .add_content("application/json", salvo::oapi::Content::new(wrapper_schema))
            .into()
    }
}

pub struct ApiOk<T>(pub Option<T>);

impl<T> ApiOk<T>
where
    T: Serialize + Send,
{
    pub fn to_response(self) -> ApiResponse<T> {
        ApiResponse {
            code: ApiResponseCode::Ok as i32,
            message: None,
            data: self.0
        } 
    }
}

impl<T> ApiOk<T>
where
    T: ToSchema,
{
    pub fn to_openapi_response(components: &mut Components) -> RefOr<salvo::oapi::Response> {
        ApiResponse::<T>::to_response(components)
    }
}

#[async_trait]
impl<T> Writer for ApiOk<T>
where
    T: Serialize + Send,
{
    async fn write(mut self, _req: &mut Request, _depot: &mut Depot, resp: &mut Response) {
        resp.render(Json(self.to_response()));
    }
}

impl<T> EndpointOutRegister for ApiOk<T>
where
    T: ToSchema,
{
    #[inline]
    fn register(components: &mut Components, operation: &mut Operation) {
        operation
            .responses
            .insert("200", Self::to_openapi_response(components));
    }
}

pub enum ApiError {
    New(ApiResponseCode, String),
    UnknownError(Option<String>),
    DbError(Option<String>),
    RequestError(Option<String>),
    AuthError(Option<String>),
    AuthorizedError(Option<String>),
}

impl ApiError {
    pub fn to_response(self) -> ApiResponse<()> {
        let (code, msg) = match self {
            ApiError::New(code, msg) => (code, msg),
            ApiError::UnknownError(msg) => (
                ApiResponseCode::UnknownError,
                msg.unwrap_or(String::from("An error occurred while processing your request. Please try again later."))
            ),
            ApiError::DbError(msg) => (
                ApiResponseCode::DbError,
                msg.unwrap_or(String::from("An error occurred while processing your request. Please try again later."))
            ),
            ApiError::RequestError(msg) => (
                ApiResponseCode::RequestError,
                msg.unwrap_or(String::from("Your request is invalid. Please check and try again."))
            ),
            ApiError::AuthError(msg) => (
                ApiResponseCode::AuthError,
                msg.unwrap_or(String::from("Unauthorized access. Please login and try again."))
            ),
            ApiError::AuthorizedError(msg) => (
                ApiResponseCode::AuthorizedError,
                msg.unwrap_or(String::from("You are not authorized to access this resource."))
            ),
        };
        ApiResponse {
            code: code as i32,
            message: Some(msg),
            data: None,
        }
    }

    pub fn to_openapi_response(components: &mut Components) -> RefOr<salvo::oapi::Response> {
        ApiResponse::<()>::to_response(components)
    }
}

#[async_trait]
impl Writer for ApiError {
    async fn write(mut self, _req: &mut Request, _depot: &mut Depot, resp: &mut Response) {
        resp.render(Json(self.to_response()));
    }
}

impl EndpointOutRegister for ApiError {
    #[inline]
    // ignore the error response, we all use status 200
    fn register(_components: &mut Components, _operation: &mut Operation) {
        // operation
        //     .responses
        //     .insert("400", Self::to_openapi_response(components));
    }
}

pub type ApiResult<T> = Result<ApiOk<T>, ApiError>;
