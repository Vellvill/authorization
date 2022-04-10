CREATE TABLE IF NOT EXISTS users(
    id bigserial not null primary key,
    email varchar not null unique,
    encrypted_password varchar not null
);
---- create above / drop below ----
DROP TABLE users;