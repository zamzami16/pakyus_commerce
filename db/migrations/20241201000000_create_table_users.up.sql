-- Add uuid extension
create extension if not exists "uuid-ossp";

create table users
(
    id         uuid         not null default uuid_generate_v4() primary key,
    username   varchar(50)  not null unique,
    name       varchar(100) not null,
    password   varchar(100) not null,
    token      varchar(100) null,
    created_at bigint       not null,
    updated_at bigint       not null
);