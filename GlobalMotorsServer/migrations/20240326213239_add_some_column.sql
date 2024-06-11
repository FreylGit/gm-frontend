-- +goose Up
create TABLE Auto(
    id serial primary key,
    title varchar(256) not null,
    price decimal not null ,
    address varchar(256) not null ,
    description text not null ,
    image_path varchar(128) not null
);

create table Detail(
    id serial primary key,
    auto_id int references Auto,
    mark varchar(256) not null ,
    model varchar(256) not null ,
    year integer not null ,
    country varchar(128) not null ,
    color varchar(128) not null ,
    body_type varchar(128) not null ,
    mileage int not null ,
    fuel varchar(128) not null ,
    type_drive varchar(128) not null ,
    gearbox_type varchar(128) not null ,
    engine_power float not null
);
create index index_detail_id_autoId
on Detail(id,auto_id);

create table Option(
    id serial primary key ,
    auto_id int references Auto,
    name varchar(128) not null ,
    value varchar(128) not null
);
create index index_option_id_autoId
    on Detail(id,auto_id);

create table Image(
    id serial primary key ,
    auto_id int references Auto,
    image_path varchar(128) not null
);
create index index_image_id_autoId
    on Detail(id,auto_id);
-- +goose Down

drop index  index_detail_id_autoId;
drop index  index_option_id_autoId;
drop index  index_image_id_autoId;

drop table Image;
drop table Option;
drop table Detail;
drop table Auto;