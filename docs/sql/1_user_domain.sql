drop table if exists gc_user;

-- nickname
-- email
-- password
-- createAt
-- deleteAt

create table gc_user
(
    id        bigserial primary key,
    nickname  varchar(15) not null,
    email     varchar(35) not null,
    password  varchar(55) not null,
    create_at timestamp   not null default now(),
    delete_at timestamp            default null
);

create unique index uix_gc_user_email ON gc_user (email);