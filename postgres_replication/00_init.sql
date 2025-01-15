CREATE USER replicator WITH REPLICATION ENCRYPTED PASSWORD '0UHl3Q14MnFL';
CREATE USER wolfgres_user PASSWORD 'K7E0Tj2y1qlq';
SELECT pg_create_physical_replication_slot('streaming_standby_server');
