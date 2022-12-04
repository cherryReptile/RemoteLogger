PRAGMA foreign_keys = ON;

CREATE TABLE users
(
    id         integer primary key autoincrement,
    tg_id      integer     not null unique,
    first_name text        not null,
    last_name  text        not null,
    username   text unique not null,
    photo_url  text        not null,
    created_at datetime    not null
);