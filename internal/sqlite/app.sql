PRAGMA foreign_keys = ON;

CREATE TABLE users
(
    id         integer primary key autoincrement,
    email      text not null unique ,
    password   text     not null,
    created_at datetime not null
);