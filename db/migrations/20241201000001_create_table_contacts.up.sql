create table contacts
(
    id         uuid         not null default uuid_generate_v4() primary key,
    first_name varchar(100) not null,
    last_name  varchar(100) null,
    email      varchar(100) null,
    phone      varchar(100) null,
    user_id    uuid         not null,
    created_at bigint       not null,
    updated_at bigint       not null,
    constraint fk_contacts_user_id foreign key (user_id) references users (id)
);