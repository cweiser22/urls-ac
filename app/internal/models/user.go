package models

type User struct {
	ID           int    `db:"id" json:"id"`
	Email        string `db:"email" json:"email"`
	PasswordHash string `db:"password_hash" json:"-"`
}

func NewUser(id int, email, passwordHash string) *User {
	return &User{
		ID:           id,
		Email:        email,
		PasswordHash: passwordHash,
	}
}
