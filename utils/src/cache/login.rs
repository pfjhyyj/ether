
// get user login token cache key
pub fn get_user_login_token_key(user_id: i64) -> String {
    format!("token:{}", user_id)
}
