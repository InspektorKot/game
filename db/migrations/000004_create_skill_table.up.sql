CREATE TABLE skill(
id serial primary key ,
class_id int not null ,
name varchar not null ,
description text ,
skill_type int not null default 0,
unlock_level int not null,
base_cooldown int,
created_at timestamp with time zone default current_timestamp not null,
updated_at timestamp with time zone default current_timestamp not null
);