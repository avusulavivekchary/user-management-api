package models

type CreateUserRequest struct {
	Name string `json:"name"`
	DOB  string `json:"dob"`
}

type UpdateUserRequest struct {
	Name string `json:"name"`
	DOB  string `json:"dob"`
}

type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age,omitempty"`
}
