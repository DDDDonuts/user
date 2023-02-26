package inmem

type InMemUserRepository struct{}

var (
	users []string
)

func NewInMemUserRepository() InMemUserRepository {
	ur := InMemUserRepository{}
	return ur
}

func (ur InMemUserRepository) CreateNewUser() {
	users = append(users, "Gustav")
}

func (ur InMemUserRepository) GetAllUsers() []string {
	return users
}
