package domain

type User struct {
	ID   int
	Name string
}

type UserRepository interface {
	Save(User) error
	FindByID(id int) (User, error)
}
