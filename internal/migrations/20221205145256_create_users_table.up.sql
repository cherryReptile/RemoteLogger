CREATE TABLE users
(
    id            bigserial                   not null primary key,
    unique_raw    varchar default '',
    password      varchar default '',
    authorized_by varchar                     not null,
    created_at    timestamp without time zone not null
);