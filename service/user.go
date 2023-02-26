package service

type userRepo interface {
	CreateNewUser()
	GetAllUsers() []string
}

type UserService struct {
	repo userRepo
}

func NewUserService(repo userRepo) UserService {
	us := UserService{repo}
	return us
}

func (us UserService) GetById() {

}

func (us UserService) Create() {
	us.repo.CreateNewUser()
}

func (us UserService) GetAll() []string {
	users := us.repo.GetAllUsers()
	return users
}
