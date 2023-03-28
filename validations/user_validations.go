package validations

import (
	"net/mail"

	"github.com/rizqyep/quicksign/domain"
	"github.com/rizqyep/quicksign/repository"
)

var repositoryInstance = repository.InitRepository()

func ValidateRegistrationRequest(user domain.User) (bool, map[string]string) {
	errors := make(map[string]string)

	if !validateUsername(user) {
		errors["username"] = "Username is already taken"
	}

	isEmailValid, emailError := validateEmail(user)

	if !isEmailValid {
		errors["email"] = emailError
	}

	return len(errors) == 0, errors
}

func ValidateLoginRequest(user domain.User) (bool, map[string]string) {
	_, err := mail.ParseAddress(user.Email)
	errors := make(map[string]string)
	if err != nil {
		errors["email"] = "E-mail should be a correct e-mail format"
	}

	return len(errors) == 0, errors
}

func validateUsername(user domain.User) bool {
	result, _ := repositoryInstance.UserRepository.GetOne(user, "username")

	return !(result != domain.User{})
}

func validateEmail(user domain.User) (bool, string) {
	_, err := mail.ParseAddress(user.Email)

	if err != nil {
		return false, "E-mail should be a correct e-mail format"
	}

	result, _ := repositoryInstance.UserRepository.GetOne(user, "email")

	if result != (domain.User{}) {
		return false, "E-mail already taken"
	}

	return true, ""
}
