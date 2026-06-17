package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
}

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
	DOB  string `json:"dob" validate:"required"`
}

type UserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age"`
}
