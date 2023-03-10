CREATE TABLE IF NOT EXISTS comments (
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
    product_id bigint NOT NULL REFERENCES products ON DELETE CASCADE,
    comment text NOT NULL,
    rating int NOT NULL DEFAULT 0,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    deleted_at timestamp(0)
    );