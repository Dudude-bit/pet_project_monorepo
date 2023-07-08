package database

type CreateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetUserDTO struct {
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
}

type GetUserReturn struct {
}

type CreateVideoReturn struct {
}

type GetVideoReturn struct {
}

type CreateDonationReturn struct {
}

type UpdateDonationStatusReturn struct {
}
