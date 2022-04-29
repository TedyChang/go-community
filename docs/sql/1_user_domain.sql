drop table if exists gc_user;

-- nickname
-- email
-- password
-- createAt
-- deleteAt

create table gc_user
(
    id         bigserial primary key,
    nickname   varchar(15) not null,
    email      varchar(35) not null,
    password   varchar(55) not null,
    created_at timestamp   not null,
    delete_at  timestamp   not null
);