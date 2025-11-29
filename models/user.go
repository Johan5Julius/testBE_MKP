package models

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserCreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserUpdateRequest struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
