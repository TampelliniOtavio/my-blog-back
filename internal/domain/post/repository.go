package post

type Repository interface {
	GetAllPosts(limit int, offset int, authUserId int64) (*[]Post, error)
	AddPost(post *Post) (*Post, error)
	GetPost(xid string, authUserId int64) (*Post, error)
	AddLikeToPost(post *Post, userId int64) error
	RemoveLikeFromPost(post *Post, userId int64) error
	DeletePost(post *Post, userId int64) error
}
