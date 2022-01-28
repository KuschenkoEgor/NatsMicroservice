CREATE OR REPLACE PROCEDURE insert_data_items(chrt_id integer, track_num varchar(50),price integer,rid varchar(50),name varchar (50), sale integer, size varchar (50), total_price integer, nm_id integer, brand varchar (50), status integer)
LANGUAGE SQL
AS $$
INSERT INTO items(chrt_id,track_number,price,rid,name,sale,size,total_price,nm_id,brand,status) VALUES (chrt_id,track_num,price,rid,name,sale,size,total_price,nm_id,brand,status);
$$;

////////////////

CREATE OR REPLACE PROCEDURE insert_data_payment(transaction varchar(50), request_id varchar(50),currency varchar(50),provider varchar(50),amount integer, payment_dt integer, bank varchar (50), delivery_cost integer, goods_total integer, custom_fee integer)
    LANGUAGE SQL
    AS $$
INSERT INTO payment(transaction,request_id,currency,provider,amount,payment_dt,bank,delivery_cost,goods_total,custom_fee) VALUES (transaction,request_id,currency,provider,amount,payment_dt,bank,delivery_cost,goods_total,custom_fee);
$$;

//////////////////

CREATE OR REPLACE PROCEDURE insert_data_delivery(name varchar (50), phone varchar (50), zip varchar (50), city varchar (50), address varchar (50), region varchar (50), email varchar (50))
    LANGUAGE SQL
    AS $$
INSERT INTO delivery(name,phone,zip,city,address,region,email) VALUES (name,phone,zip,city,address,region,email);
$$;

/////////////////

CREATE OR REPLACE PROCEDURE insert_data_main(delivery_id integer , payment_id integer , items_id integer , order_uid varchar (50),track_number varchar (50),entry varchar (50),locale varchar (50),internal_signature varchar (50), customer_id varchar (50), delivery_service varchar (50),shardkey varchar (10), sm_id integer, date_created varchar (50), oof_shard varchar (10))
    LANGUAGE SQL
    AS $$
INSERT INTO main(delivery_id,payment_id,items_id,order_uid,track_number,entry,locale,internal_signature,customer_id,delivery_service,shardkey,sm_id,date_created,oof_shard) VALUES (delivery_id,payment_id,items_id,order_uid,track_number,entry,locale,internal_signature,customer_id,delivery_service,shardkey,sm_id,date_created,oof_shard);
$$;