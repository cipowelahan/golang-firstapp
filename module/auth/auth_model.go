package auth

type AuthToken struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

type AuthLogin struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type AuthRegister struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6,max=12"`
}
