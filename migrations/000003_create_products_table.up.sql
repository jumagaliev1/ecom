CREATE TABLE IF NOT EXISTS products (
    id bigserial PRIMARY KEY,
    category_id bigint NOT NULL REFERENCES categories ON DELETE CASCADE,
    user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
    title text NOT NULL,
    description text NOT NULL,
    price bigint NOT NULL,
    rating float NOT NULL DEFAULT 0,
    stock int NOT NULL DEFAULT 1,
    images text[],
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    deleted_at timestamp(0)
);