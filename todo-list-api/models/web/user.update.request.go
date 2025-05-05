package web

type UserUpdateRequest struct {
	Id       int
	Username string
	Email    string
	Password string
}

type UserUpdateRequestUsername struct {
	Id       int
	Username string
}
