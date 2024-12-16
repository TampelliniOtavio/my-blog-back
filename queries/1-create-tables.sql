CREATE SCHEMA my_blog;

CREATE TABLE my_blog.users (
    id SERIAL PRIMARY KEY,
    xid char(20) NOT NULL,
    username varchar(20) NOT NULL,
    email varchar(20) NOT NULL,
    password varchar(255) NOT NULL
);

CREATE UNIQUE INDEX users_xid ON my_blog.users(xid);
CREATE UNIQUE INDEX users_username ON my_blog.users(username);
CREATE UNIQUE INDEX users_email ON my_blog.users(email);

CREATE TABLE my_blog.posts (
	xid bpchar(20) NOT NULL primary key,
	post varchar(128) NOT NULL,
	created_by int NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL
)


ALTER TABLE my_blog.posts ADD CONSTRAINT posts_users_fk FOREIGN KEY (created_by) REFERENCES my_blog.users(id) ON DELETE CASCADE ON UPDATE CASCADE;
CREATE INDEX posts_created_at_idx ON my_blog.posts (created_at DESC);
CREATE INDEX posts_updated_at_idx ON my_blog.posts (updated_at DESC);
CREATE INDEX posts_created_by_idx ON my_blog.posts (created_by);
