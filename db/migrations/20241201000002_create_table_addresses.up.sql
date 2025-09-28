create table addresses
(
    id           uuid         not null default uuid_generate_v4() primary key,
    contact_id   uuid         not null,
    street       varchar(255) null,
    city         varchar(255) null,
    province     varchar(255) null,
    postal_code  varchar(10)  null,
    country      varchar(100) null,
    created_at   bigint       not null,
    updated_at   bigint       not null,
    constraint fk_addresses_contact_id foreign key (contact_id) references contacts (id)
);