INSERT INTO wfg.account (customer_id, account_type_id, balace) VALUES 
    ((SELECT customer_id FROM wfg.customer ORDER BY RANDOM() LIMIT 1),
     (floor(random() * 3) + 1),
      (floor(random() * 100) + 1)::double precision);
