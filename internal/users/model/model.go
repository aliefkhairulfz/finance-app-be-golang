package model

type User struct {
	ID           int     `json:"id"`
	Email        string  `json:"email"`
	Password     *string `json:"password"`
	Avatar       *string `json:"avatar"`
	IsActive     bool    `json:"is_active"`
	RefreshToken *string `json:"refresh_token"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type UserResponse struct {
	ID        int     `db:"id"`
	Email     string  `db:"email"`
	Avatar    *string `db:"avatar"`
	IsActive  bool    `db:"is_active"`
	CreatedAt string  `db:"created_at"`
	UpdatedAt string  `db:"updated_at"`
}
