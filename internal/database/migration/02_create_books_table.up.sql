CREATE TABLE IF NOT EXISTS go_basic_db.books (
    id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    author varchar(255) NOT NULL,
    name varchar(255)
);