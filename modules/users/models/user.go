package users_models

type Role string

const (
	Admin Role = "ADMIN"
	Basic Role = "BASIC"
)

type User struct {
	ID       string `json:"id"`
	Active   bool   `json:"active"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}
