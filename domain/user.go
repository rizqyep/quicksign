package domain

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u *User) HashPassword() {
	bytePassword := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.Password = string(hashedPassword)
}

func (u *User) HidePassword() {
	u.Password = ""
}
