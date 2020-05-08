CREATE TABLE player (
id serial primary key ,
name varchar not null unique ,
class_id int not null ,
level int not null ,
current_exp int not null ,
level_exp int not null,
min_damage int not null,
max_damage int not null ,
health int not null ,
max_health int not null ,
critical_chance int not null ,
created_at timestamp with time zone default current_timestamp not null,
updated_at timestamp with time zone default current_timestamp not null
);