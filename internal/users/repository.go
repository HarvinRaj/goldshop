package users

type Repository interface {
	Save(user Users) error
}

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u UserRepository) Save(user Users) error {
	return nil
}
