package post

type Repository interface {
	GetAllPosts(limit int, offset int) (*[]Post, error)
}
