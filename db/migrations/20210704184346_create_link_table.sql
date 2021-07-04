-- +goose Up
-- +goose StatementBegin
CREATE TABLE link (
    id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    type INT NOT NULL,
    created_date DATETIME NOT NULL,
    user_id BIGINT NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES user(id)
);
-- +goose StatementEnd
