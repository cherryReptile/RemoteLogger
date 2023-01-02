create table user_profiles
(
    id         bigserial not null primary key,
    first_name varchar,
    last_name  varchar,
    address    varchar,
    other_data json    default '{}',
    user_id    varchar references users (id) on delete cascade,
    created_at timestamp without time zone not null
)