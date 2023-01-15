create table users_providers_data
(
    id          bigserial                   not null primary key,
    user_data   json                        not null,
    user_id     varchar references users (id) on delete cascade,
    provider_id bigserial references providers (id) on delete cascade,
    username    varchar                     not null,
    created_at  timestamp without time zone not null
)