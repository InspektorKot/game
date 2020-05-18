CREATE TABLE enemy(
id serial primary key ,
name varchar not null ,
level int not null ,
exp int not null ,
min_damage int not null ,
max_damage int not null ,
health int not null,
created_at timestamp with time zone default current_timestamp not null,
updated_at timestamp with time zone default current_timestamp not null
);