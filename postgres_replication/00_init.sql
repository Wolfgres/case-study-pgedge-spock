create user replicator with replication encrypted password 'replicator_password';
create user wolfgres_user password 'password';
create database wolfgres_db;
select pg_create_physical_replication_slot('streaming_standby_server');
