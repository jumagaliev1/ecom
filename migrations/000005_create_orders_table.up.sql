CREATE TABLE IF NOT EXISTS orders (
    id bigserial PRIMARY KEY,
    order_status text NOT NULL DEFAULT 'CREATED',
    cart_id bigint NOT NULL REFERENCES carts ON DELETE CASCADE,
    quantity int NOT NULL DEFAULT 1,
    total_price bigint NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    deleted_at timestamp(0)
);