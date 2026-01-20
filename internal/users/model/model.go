package model

type User struct {
	ID           int     `json:"id"`
	Email        string  `json:"email"`
	Password     *string `json:"password"`
	Avatar       string  `json:"avatar"`
	IsActive     bool    `json:"is_active"`
	RefreshToken *string `json:"refresh_token"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}
