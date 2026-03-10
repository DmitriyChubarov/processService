CREATE TABLE IF NOT EXISTS processing (
    id SERIAL PRIMARY KEY,
    user_id BIGINT,
    original_file_path TEXT,
    processed_file_path TEXT,
    upload_datetime BIGINT,
    status TEXT
);