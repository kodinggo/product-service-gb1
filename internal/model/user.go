package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
