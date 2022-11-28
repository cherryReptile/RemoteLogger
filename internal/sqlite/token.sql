CREATE TABLE tokens
(
    id         integer primary key autoincrement,
    token      text    not null unique,
    user_id    integer not null,
    created_at datetime not null,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);