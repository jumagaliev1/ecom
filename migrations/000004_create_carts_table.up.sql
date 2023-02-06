CREATE TABLE IF NOT EXISTS carts (
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
    product_id bigint NOT NULL REFERENCES products ON DELETE CASCADE,
    quantity int NOT NULL DEFAULT 1
    );