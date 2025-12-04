package adminrequest

type RegisterAdminRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginAdminRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UpdateProfileRequest struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}