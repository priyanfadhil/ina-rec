package user

type User struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt string `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
}

func (User) ModelName() string {
	return "users"
}

func (User) TableName() string {
	return "public.users"
}

type GetUserReq struct {
	Id int `uri:"id"`
}

type CreateUserRequest struct {
	ID        int    `json:"-" gorm:"primaryKey" uri:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
}

func (CreateUserRequest) ModelName() string {
	return "users"
}

func (CreateUserRequest) TableName() string {
	return "public.users"
}

type UpdateUserRequest struct {
	ID        int    `json:"-" gorm:"primaryKey" uri:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	UpdatedAt string `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
}

func (UpdateUserRequest) ModelName() string {
	return "users"
}

func (UpdateUserRequest) TableName() string {
	return "public.users"
}
