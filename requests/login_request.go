package requests

import "strings"

type LoginRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (l *LoginRequest) Validate() map[string]string {
	errors := make(map[string]string)

	if strings.TrimSpace(l.Email) == "" {
		errors["email"] = "Email can't be blank"
	}

	if strings.TrimSpace(l.Password) == "" {
		errors["password"] = "Password can't be blank"
	}

	return errors
}
