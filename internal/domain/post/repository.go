package post

type Repository interface {
	GetAllPosts(limit int, offset int) (*[]Post, error)
	AddPost(post *Post) (*Post, error)
	GetPost(xid string) (*Post, error)
	AddLikeToPost(post *Post, userId int64) error
	RemoveLikeFromPost(post *Post, userId int64) error
	DeletePost(post *Post, userId int64) error
}
