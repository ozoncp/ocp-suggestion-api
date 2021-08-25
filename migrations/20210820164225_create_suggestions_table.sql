-- +goose Up
-- +goose StatementBegin
CREATE TABLE suggestions
(
    id        BIGSERIAL UNIQUE PRIMARY KEY,
    user_id   BIGINT NOT NULL,
    course_id BIGINT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE suggestions;
-- +goose StatementEnd
