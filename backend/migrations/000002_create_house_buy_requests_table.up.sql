

CREATE TABLE house_buy_requests (
    id uuid PRIMARY KEY DEFAULT uuidv7(),

    email TEXT NOT NULL,

    first_name TEXT NOT NULL,

    last_name TEXT NOT NULL,

    address TEXT NOT NULL,

    was_offered BOOLEAN DEFAULT false,

    time_received TIMESTAMPTZ DEFAULT now()
);