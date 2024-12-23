pub fn parse_page_and_size(page: Option<u64>, size: Option<u64>) -> (u64, u64) {
    let mut offset: u64 = 0;
    let mut limit: u64 = 20;

    if let Some(page) = page {
        if page < 1 {
            offset = 0;
        } else {
            offset = (page - 1) as u64 * limit;
        }
    }

    if let Some(size) = size {
        if size < 1 {
            limit = 20;
        } else if size > 100 {
            limit = 100;
        } else {
            limit = size as u64;
        }
    }

    (offset, limit)
}