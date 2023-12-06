package postings

type UserService interface {
	GetUserById(userID uint) (User, error)
}

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Image    string `json:"foto_profil"`
}
