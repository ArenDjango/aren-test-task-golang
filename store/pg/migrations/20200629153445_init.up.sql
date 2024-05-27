CREATE TABLE rates (
	id uuid NOT NULL,
    ask_price float NOT NULL,
    bid_price float NOT NULL,
    time_stamp timestamp default current_timestamp,
	CONSTRAINT "pk_user_id" PRIMARY KEY (id)
);