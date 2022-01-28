CREATE TABLE delivery (
    delivery_id SERIAL PRIMARY KEY,
    name VARCHAR (30),
    phone VARCHAR (30),
    zip VARCHAR (30),
    city VARCHAR (30),
    address VARCHAR (30),
    region VARCHAR (30),
    email VARCHAR (30)
);

CREATE TABLE payment (
    payment_id SERIAL PRIMARY KEY,
    transaction VARCHAR (30),
    request_id VARCHAR (30),
    currency VARCHAR (10),
    provider VARCHAR (30),
    amount INT,
    payment_dt INT,
    bank VARCHAR (30),
    delivery_cost INT,
    goods_total INT,
    custom_fee INT
);

CREATE TABLE items (
    items_id SERIAL PRIMARY KEY,
    chrt_id INT,
    track_number VARCHAR (30),
    price INT,
    rid VARCHAR (30),
    name VARCHAR (30),
    sale VARCHAR (30),
    size VARCHAR (30),
    total_price INT,
    nm_id INT,
    brand VARCHAR (30),
    status INT
);
CREATE TABLE main (
                      main_id SERIAL PRIMARY KEY,
                      order_uid VARCHAR (30),
                      track_number VARCHAR (30),
                      entry VARCHAR (30),
                      delivery_id INT,
                      items_id INT,
                      payment_id INT,
                      locale VARCHAR (30),
                      internal_signature VARCHAR (30),
                      customer_id VARCHAR (30),
                      delivery_service VARCHAR (30),
                      shardkey VARCHAR (30),
                      sm_id INT,
                      date_created VARCHAR (50),
                      oof_shard VARCHAR (10),
                      FOREIGN KEY (payment_id) REFERENCES payment (payment_id),
                      FOREIGN KEY (items_id) REFERENCES items (items_id),
                      FOREIGN KEY (delivery_id) REFERENCES delivery (delivery_id)
);