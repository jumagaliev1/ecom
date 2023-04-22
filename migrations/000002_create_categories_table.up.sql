CREATE TABLE IF NOT EXISTS categories (
    id bigserial PRIMARY KEY,
    title text NOT NULL
);

INSERT INTO categories (title) VALUES ("Food");
INSERT INTO categories (title) VALUES ("Digital");
INSERT INTO categories (title) VALUES ("Furniture");
INSERT INTO categories (title) VALUES ("Health");
INSERT INTO categories (title) VALUES ("Other");
