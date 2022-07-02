package dtos

type SignupDto struct {
	FullName   *string `json:"full_name" validate:"required,min=2,max=20"`
	Email      *string `json:"email" validate:"required,email"`
	Password   *string `json:"password" validate:"required,min=6"`
	ProfilePic *string `json:"profile_pic" validate:"required"`
}
