

ALTER TABLE house_buy_requests
    ADD COLUMN comments TEXT CHECK (LENGTH(comments) <= 200);