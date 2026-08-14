

CREATE TABLE users (
    id uuid PRIMARY KEY default uuidv7(),

    email TEXT UNIQUE NOT NULL,

    time_created TIMESTAMPTZ NOT NULL DEFAULT now()
);

ALTER TABLE house_buy_requests
    DROP COLUMN email,
    ADD COLUMN user_id uuid NOT NULL REFERENCES users(id);