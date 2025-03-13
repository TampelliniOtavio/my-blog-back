package user

type Service interface {
	GetByUsername(username string) (*User, error)
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) GetByUsername(username string) (*User, error) {
	return s.Repository.GetByUsername(username)
}
