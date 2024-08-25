CREATE TABLE rates (
	id uuid NOT NULL,
    ask_price float NOT NULL,
    bid_price float NOT NULL,
    time_stamp timestamp default current_timestamp,
	CONSTRAINT "pk_user_id" PRIMARY KEY (id)
);

CREATE TABLE outbox (
    id SERIAL PRIMARY KEY,
    message_type VARCHAR(50),
    unique_id string,
    chronometric_id VARCHAR(50),
    payload JSONB,
    status VARCHAR(20),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    processed_at TIMESTAMPTZ
);