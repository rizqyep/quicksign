package domain

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

func (u *User) ValidateRegistrationRequest() (bool, map[string]string) {
	errors := make(map[string]string)

	if len(errors) != 0 {
		return false, errors
	}

	return true, errors
}
