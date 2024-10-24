package user

type UpdateUserRequest struct {
	Id        int32  `json:"id" validate:"required" param:"id"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required,min=1,max=32"`
	LastName  string `json:"last_name" validate:"required,min=1,max=32"`
	Avatar    string `json:"avatar"`
}
