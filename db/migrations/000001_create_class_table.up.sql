CREATE TABLE class (
id serial primary key ,
name varchar not null unique ,
min_damage int not null,
max_damage int not null ,
health int not null ,
critical_chance int not null ,
created_at timestamp with time zone default current_timestamp not null,
updated_at timestamp with time zone default current_timestamp not null
);