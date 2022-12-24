create table users_auth_providers_data
(
    id          bigserial                   not null primary key,
    user_data   json                        not null,
    user_id     varchar references users (id) on delete cascade,
    provider_id bigserial references auth_providers (id) on delete cascade,
    created_at  timestamp without time zone not null
)