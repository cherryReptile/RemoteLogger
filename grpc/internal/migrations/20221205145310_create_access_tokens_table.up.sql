create table access_tokens
(
    id         bigserial primary key not null,
    token      text    not null unique,
    user_id    bigserial references users(id) on delete cascade,
    created_at timestamp without time zone not null
);