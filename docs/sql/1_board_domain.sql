drop table if exists gc_board;

-- title
-- content
-- board_category_id
-- user_id
-- createAt
-- deleteAt

create table gc_board
(
    id                   bigserial primary key,
    title                varchar(40)  not null,
    content              varchar(255) not null,
    gc_board_category_id bigint       not null,
    gc_user_id           bigint       not null,
    create_at            timestamp    not null default now(),
    delete_at            timestamp             default null
);

create unique index idx_gc_board_title ON gc_board (title);