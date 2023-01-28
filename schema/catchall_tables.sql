-- DROP DATABASE IF EXISTS "catchall-db";
-- CREATE DATABASE "catchall-db";

CREATE TABLE IF NOT EXISTS domain_names (
   id serial primary key,
   name VARCHAR (255) not null unique,
   deliveredevent int not null,
   bouncedevent int not null,
   status VARCHAR (50) not null
);

create unique index domain_names_unique_idx on domain_names (lower(name));
