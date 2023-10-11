package user

type GetUser struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (GetUserResponse) ModelName() string {
	return "users"
}

func (GetUserResponse) TableName() string {
	return "public.users"
}
