

ALTER TABLE house_buy_requests
    ADD CONSTRAINT email_not_empty CHECK (length(trim(email)) > 0),
    ADD CONSTRAINT first_name_not_empty CHECK (length(trim(first_name)) > 0),
    ADD CONSTRAINT last_name_not_empty CHECK (length(trim(last_name)) > 0),
    ADD CONSTRAINT address_not_empty CHECK (length(trim(address)) > 0);