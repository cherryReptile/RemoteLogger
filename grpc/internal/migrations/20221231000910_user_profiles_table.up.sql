create table user_profiles
(
    id         bigserial not null primary key,
    first_name varchar   not null,
    last_name  varchar   not null,
    address    varchar   not null,
    other_data json default null,
    user_id    varchar references users (id) on delete cascade,
    created_at timestamp without time zone
)