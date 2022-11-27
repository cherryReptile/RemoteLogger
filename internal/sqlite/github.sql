PRAGMA foreign_keys = ON;

CREATE TABLE users
(
    id         integer primary key autoincrement,
    login      text     not null unique,
    email      text default '',
    avatar_url text     not null,
    created_at datetime not null
);