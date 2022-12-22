CREATE TABLE users
(
    id         varchar                     primary key,
    login      varchar                     not null,
    password   varchar default '',
    created_at timestamp without time zone not null
);