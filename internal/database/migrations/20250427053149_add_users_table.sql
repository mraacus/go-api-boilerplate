-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id    BIGSERIAL PRIMARY KEY,
  name  text      NOT NULL,
  role  text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE users;
-- +goose StatementEnd
