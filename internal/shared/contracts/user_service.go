package contracts

type User struct {
	ID    uint
	Email string
}

type UserService interface {
	GetUser(id uint) (User, error)
}
