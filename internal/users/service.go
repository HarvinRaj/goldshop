package users

type Service interface {
	CreateUser(users *Users) (*Users, error)
}

type UserService struct {
	repo Repository
}

func NewUserService(repo Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u UserService) CreateUser(*Users) (*Users, error) {
	return &Users{}, nil
}
