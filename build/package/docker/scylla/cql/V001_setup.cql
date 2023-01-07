CREATE KEYSPACE account_consumer_service WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

CREATE TABLE account_consumer_service.account
(
	email text
	,name text
	,alias text
	,city text
	,district text
	,public_place text
	,zip_code text
	,full_number text
	,status text
	,date_time text
	,PRIMARY KEY(email, full_number)
);
