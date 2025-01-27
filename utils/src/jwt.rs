use jsonwebtoken::{decode, encode, errors::Error, DecodingKey, EncodingKey, Header, Validation};
use serde::{de::DeserializeOwned, Serialize};

pub fn generate_jwt_token<T>(claims: &T) -> Result<String, Error>
where
    T: Serialize,
{
    let jwt_secret = crate::config::global().get_string("app.secret").unwrap_or_default();
    let key = EncodingKey::from_secret(jwt_secret.as_bytes());
    let token = encode(&Header::default(), &claims, &key);

    token
}

pub fn verify_jwt_token<T>(token: &str) -> Result<T, Error>
where
    T: DeserializeOwned,
{
    let jwt_secret = crate::config::global().get_string("app.secret").unwrap_or_default();
    let key = DecodingKey::from_secret(jwt_secret.as_bytes());
    let token_data = decode::<T>(&token, &key, &Validation::default()).map(|data| data.claims);

    token_data
}
