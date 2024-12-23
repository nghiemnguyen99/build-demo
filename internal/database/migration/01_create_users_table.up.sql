CREATE TABLE IF NOT EXISTS go_basic_db.users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    email varchar(255) NOT NULL,
    pass varchar(100)
);