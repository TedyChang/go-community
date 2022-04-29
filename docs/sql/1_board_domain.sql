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
    created_at           timestamp    not null,
    delete_at            timestamp    not null
);