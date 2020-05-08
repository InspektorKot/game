CREATE TABLE player_item(
id serial primary key ,
player_id int not null ,
item_id int not null ,
amount int not null,
created_at timestamp with time zone default current_timestamp not null,
updated_at timestamp with time zone default current_timestamp not null
);