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
	}

	if strings.TrimSpace(s.Email) == "" {
		errors["email"] = "Email can't be blank"
	}

	if strings.TrimSpace(s.Password) == "" {
		errors["password"] = "Password can't be blank"
	}

	if !utils.IsValidEmail(s.Email) {
		errors["email"] = "Email is invalid"
	}

	return errors
}
