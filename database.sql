/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */

CREATE TABLE users (
	id bigserial PRIMARY KEY,
	full_name varchar(60) NOT NULL,
	password varchar(64) NOT NULL,
	country_code varchar(7) NOT NULL,
	phone_number varchar(13) UNIQUE NOT NULL,
	successful_login integer DEFAULT 0,
	created_at timestamptz NOT NULL DEFAULT (now()),
	updated_at timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON users (phone_number);
