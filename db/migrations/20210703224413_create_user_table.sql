-- +goose Up
-- +goose StatementBegin
CREATE TABLE user (
    id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT
);
-- +goose StatementEnd
