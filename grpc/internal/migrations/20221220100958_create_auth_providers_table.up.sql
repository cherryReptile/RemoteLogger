create table auth_providers
(
    id           bigserial not null primary key,
    provider     varchar   not null unique,
    unique_key   varchar   not null
)