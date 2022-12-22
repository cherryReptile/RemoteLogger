CREATE TABLE intermediate
(
    id          bigserial                   not null primary key,
    user_id     varchar references users (id) on delete cascade,
    provider_id bigserial                   not null,
    created_at  timestamp without time zone not null
);