--Execute only if you do not have a bigint data type 
--ALTER TYPE
--ALTER TABLE wfg.account ALTER COLUMN account_id TYPE bigint;
--ALTER TABLE wfg.account_type ALTER COLUMN account_type_id TYPE bigint;
--ALTER TABLE wfg.customer ALTER COLUMN customer_id TYPE bigint;
--ALTER TABLE wfg.operation ALTER COLUMN operation_id TYPE bigint;
--ALTER TABLE wfg.transaction ALTER COLUMN transaction_id TYPE bigint;

-- ALTER TABLE
ALTER TABLE wfg.account ALTER COLUMN account_id SET DEFAULT snowflake.nextval('account_account_id_seq');
ALTER TABLE wfg.account_type ALTER COLUMN account_type_id SET DEFAULT snowflake.nextval('account_type_account_type_id_seq');
--ALTER TABLE wfg.customer ALTER COLUMN customer_id SET DEFAULT snowflake.nextval();
ALTER TABLE wfg.operation ALTER COLUMN operation_id SET DEFAULT snowflake.nextval('operation_operation_id_seq');
ALTER TABLE wfg.transaction ALTER COLUMN transaction_id SET DEFAULT snowflake.nextval('transaction_transaction_id_seq');
