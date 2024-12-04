CREATE SCHEMA my_blog;

CREATE TABLE my_blog.users (
    id SERIAL PRIMARY KEY,
    xid char(20) NOT NULL,
    username varchar(20) NOT NULL,
    email varchar(20) NOT NULL,
    password varchar(255) NOT NULL
);

CREATE UNIQUE INDEX users_xid ON users(xid);
CREATE UNIQUE INDEX users_username ON users(username);
CREATE UNIQUE INDEX users_email ON users(email);
