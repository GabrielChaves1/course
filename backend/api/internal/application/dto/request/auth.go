package request

type SignInDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpDTO struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
