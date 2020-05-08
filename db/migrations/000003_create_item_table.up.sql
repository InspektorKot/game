CREATE TABLE item (
id serial primary key ,
name varchar not null,
description text,
created_at timestamp with time zone default current_timestamp not null,
updated_at timestamp with time zone default current_timestamp not null
);