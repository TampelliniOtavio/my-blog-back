alter table my_blog.posts add column deleted_at timestamp default null;

create table my_blog.likes_post (
	user_id int not null,
	post_xid bpchar(20) not null,
	created_at timestamp not null default (now())
);

alter table my_blog.likes_post add constraint likes_post_one_user_per_post unique (user_id,post_xid);

alter table my_blog.likes_post add constraint likes_post_posts_fk foreign key (post_xid) references my_blog.posts(xid) on delete cascade on update restrict;
alter table my_blog.likes_post add constraint likes_post_users_fk foreign key (user_id) references my_blog.users(id) on delete cascade on update restrict;

alter table my_blog.posts add column like_count int not null default 0;
