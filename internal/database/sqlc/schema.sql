CREATE TABLE users (
  id    BIGSERIAL PRIMARY KEY,
  name  text      NOT NULL,
  role  text
);

CREATE TABLE cards (
    id          BIGSERIAL PRIMARY KEY,
    owner_id    BIGSERIAL NOT NULL,
    type        text NOT NULL,
    number      text NOT NULL,
    exp_date    date NOT NULL,
    cvv         text NOT NULL,
    balance     numeric NOT NULL,
    created_at  timestamp DEFAULT now(),
    updated_at  timestamp DEFAULT now(),


    FOREIGN KEY (owner_id) REFERENCES users(id)
)