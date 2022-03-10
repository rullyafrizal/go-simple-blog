package requests

import (
	"strings"

	"github.com/rullyafrizal/go-simple-blog/utils"
)

type StoreUserRequest struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (s *StoreUserRequest) Validate() map[string]string {
	var errors map[string]string = make(map[string]string)

	if strings.TrimSpace(s.Name) == "" {
		errors["name"] = "Name can't be blank"
	} else {
		if len(s.Name) > 255 {
			errors["name"] = "Name can't be longer than 255 characters"
		}
	}

	if strings.TrimSpace(s.Email) == "" {
		errors["email"] = "Email can't be blank"
	} else {
		if len(s.Email) > 255 {
			errors["email"] = "Email can't be longer than 255 characters"
		}
	}

	if strings.TrimSpace(s.Password) == "" {
		errors["password"] = "Password can't be blank"
	} else {
		if len(s.Password) > 20 {
			errors["password"] = "Password can't be longer than 20 characters"
		}

		// password can't contain any whitespaace
		if strings.Contains(s.Password, " ") {
			errors["password"] = "Password can't contain any whitespace"
		}
	}

	if !utils.IsValidEmail(s.Email) {
		errors["email"] = "Email is invalid"
	}

	return errors
}
