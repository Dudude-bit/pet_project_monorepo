package user

type RegisterUserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetUserDTO struct {
	Id string `json:"id"`
}

type RegisterUserReturn struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type LoginUserReturn struct {
	AccessToken string `json:"access_token"`
}

type GetUserReturn struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
