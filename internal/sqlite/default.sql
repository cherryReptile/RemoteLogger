PRAGMA foreign_keys = ON;

CREATE TABLE users
(
    id         integer primary key autoincrement,
    login      text     not null unique,
    email      text default '',
    avatar_url text     not null,
    created_at datetime not null
);

CREATE TABLE tokens
(
    id         integer primary key autoincrement,
    token      text    not null unique,
    user_id    integer not null,
    created_at datetime not null,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);