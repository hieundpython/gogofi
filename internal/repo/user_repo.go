package repo

type UserRepo struct {
	UserId    int    `json:"user_id" db:"user_id"`
	UserName  string `json:"user_name" db:"user_name"`
	UserEmail string `json:"user_email" db:"user_email"`
}
