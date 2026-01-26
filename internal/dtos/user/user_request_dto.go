package user

type CreateUserRequest struct {
	Name     string `valid:"required,alpha,length(2|50)"`
	Email    string `valid:"required,email"`
	Password string `valid:"required,length(6|125)"`
}

type UpdateUserRoleRequest struct {
	Role string `json:"role" valid:",length(0|5)"`
}
