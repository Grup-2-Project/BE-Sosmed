package handler

type RegisterRequest struct {
	FirstName     string `json:"firstname" form:"firstname"`
	LastName    string `json:"lastname" form:"lastname"`
	Gender string `json:"gender" form:"gender"`
	Hp string `json:"hp" form:"hp"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterResponse struct {
	ID    uint   `json:"id"`
	FirstName     string `json:"firstname"`
	LastName    string `json:"lastname"`
	Gender string `json:"gender"`
	Hp string `json:"hp"`
	Email string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    uint   `json:"id"`
	FirstName     string `json:"firstname"`
	LastName    string `json:"lastname"`
	Gender string `json:"gender"`
	Hp string `json:"hp"`
	Email string `json:"email"`
	Image string `json:"image"`
	Token string `json:"token"`
}
