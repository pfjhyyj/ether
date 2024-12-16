use std::fmt::{self, Display};

use serde::{Deserialize, Serialize};

use crate::jwt::verify_jwt_token;



#[derive(Debug, Default, Clone, Serialize, Deserialize)]
pub struct Identity {
    // sub: subject, user id
    pub sub: i64,
    // exp: expiration time
    pub exp: i64,
}

impl Identity {
    pub fn new(user_id: i64, exp: i64) -> Self {
        Identity {
            sub: user_id,
            exp,
        }
    }

    pub fn empty() -> Self {
        Identity {
            sub: 0,
            exp: 0,
        }
    }

    pub fn is_valid(&self) -> bool {
        self.sub > 0 && self.exp > chrono::Utc::now().timestamp()
    }

    pub fn from_auth_token(token: String) -> Self {
        if token.len() == 0 {
            return Identity::empty();
        }
        let token = token.replace("Bearer ", "");
        let token_data = verify_jwt_token::<Identity>(&token);
        match token_data {
            Ok(data) => data,
            Err(_) => Identity::empty(),
        }
    }
}

impl Display for Identity {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        if self.sub == 0 {
            return write!(f, "<none>");
        }
        write!(f, "id:{}", self.sub)
    }
}
