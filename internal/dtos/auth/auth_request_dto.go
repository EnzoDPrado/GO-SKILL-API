package auth

type AuthUserRequestDto struct {
	Email    string `valid:"required,email"`
	Password string `valid:"required,length(6|125)"`
}
