

ALTER TABLE house_buy_requests
    DROP COLUMN was_offered,
    ADD COLUMN price_offered_cents BIGINT CHECK (price_offered_cents >= 0);