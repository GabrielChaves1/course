package request

type SignInDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResendConfirmationCodeDTO struct {
	Email string `json:"email"`
}

type SignUpDTO struct {
	Email           string `json:"email"`
	PhoneNumber     string `json:"phone_number"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type ConfirmSignUpDTO struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type ForgotPasswordDTO struct {
	Email string `json:"email"`
}

type ResetPasswordDTO struct {
	Email    string `json:"email"`
	Code     string `json:"code"`
	Password string `json:"password"`
}
