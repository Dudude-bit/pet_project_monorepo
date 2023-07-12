package database

type CreateUserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserDTO struct {
	Id string `json:"id"`
}

type GetUserByUsernameDTO struct {
	Username string `json:"username"`
}

type CreateVideoDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	User        string `json:"user"`
}

type GetVideoDTO struct {
	Id string `json:"id"`
}

type CreateDonationDTO struct {
}

type UpdateDonationStatusDTO struct {
}

type CreateUserReturn struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type GetUserReturn struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type GetUserByUsernameReturn struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateVideoReturn struct {
}

type GetVideoReturn struct {
}

type CreateDonationReturn struct {
}

type UpdateDonationStatusReturn struct {
}
