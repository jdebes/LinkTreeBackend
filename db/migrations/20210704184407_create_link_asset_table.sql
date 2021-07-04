-- +goose Up
-- +goose StatementBegin
CREATE TABLE link_asset (
    id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    type INT NOT NULL,
    name VARCHAR(144) NOT NULL,
    url VARCHAR(2048) NOT NULL,
    link_id BIGINT NOT NULL,
    platform_id BIGINT,
    CONSTRAINT fk_link_id FOREIGN KEY (link_id) REFERENCES link(id),
    CONSTRAINT fk_platform_id FOREIGN KEY (platform_id) REFERENCES platform(id)
);
-- +goose StatementEnd
