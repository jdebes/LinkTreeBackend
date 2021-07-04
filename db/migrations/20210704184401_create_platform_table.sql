-- +goose Up
-- +goose StatementBegin
CREATE TABLE platform (
    id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(144) NOT NULL,
    url VARCHAR(2048) NOT NULL,
    logo_url VARCHAR(2048) NOT NULL
);
-- +goose StatementEnd
