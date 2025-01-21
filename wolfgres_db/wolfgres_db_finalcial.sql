
-- Create Schema
CREATE SCHEMA wfg;

ALTER SCHEMA wfg OWNER TO wolfgres_user;

CREATE TABLE wfg.account (
    account_id integer NOT NULL,
    customer_id integer,
    account_type_id integer,
    balace double precision
);


ALTER TABLE wfg.account OWNER TO wolfgres_user;


CREATE SEQUENCE wfg.account_account_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE wfg.account_account_id_seq OWNER TO wolfgres_user;

ALTER SEQUENCE wfg.account_account_id_seq OWNED BY wfg.account.account_id;

CREATE TABLE wfg.account_type (
    account_type_id integer NOT NULL,
    name character varying(50),
    description text
);

ALTER TABLE wfg.account_type OWNER TO admin;

CREATE SEQUENCE wfg.account_type_account_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE wfg.account_type_account_type_id_seq OWNER TO admin;