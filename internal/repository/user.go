package repository

type UserRepository struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) GetUserById() string {
	return "vegetarian23"
}