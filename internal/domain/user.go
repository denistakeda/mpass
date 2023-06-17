package domain

type User struct {
	Login        string `db:"login"`
	PasswordHash string `db:"password"`
}
