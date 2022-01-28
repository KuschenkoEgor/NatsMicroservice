INSERT INTO main (track_number) VALUES ('WBILMTESTTRACK');
INSERT INTO payment (currency) VALUES ('USD');
INSERT INTO items (rid) VALUES ('ab4219087a764ae0btest');

SELECT currency, track_number
FROM
    main INNER JOIN payment ON main.payment_id = payment.payment_id;


SELECT *
FROM
    main INNER JOIN payment ON main.payment_id = payment.payment_id
         INNER JOIN delivery ON main.delivery_id = delivery.delivery_id
         INNER JOIN items ON main.items_id = items.items_id;

SELECT delivery_id, payment_id, items_id
FROM delivery,payment,items;
