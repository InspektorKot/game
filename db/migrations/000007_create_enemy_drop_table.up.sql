CREATE TABLE enemy_drop(
id serial primary key ,
enemy_id int not null ,
item_id int not null ,
amount int not null ,
chance float not null
);