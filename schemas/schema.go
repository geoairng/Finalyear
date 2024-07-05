package schemas

import (
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	CompanyName     string `json:"Firstname" binding:"required" example:"Afeez"`
	Email           string `json:"Email" binding:"required" example:"Lawal@gmail.com"`
	Password        string `json:"Password" binding:"required" example:"secret"`
	ConfirmPassword string `json:"ConfirmPassword" binding:"required" example:"secret"`
}

func (r *RegisterInput) VerifyPassword() bool {

	password := strings.TrimSpace(r.Password)
	confirm := strings.TrimSpace(r.ConfirmPassword)
	r.Password = password
	r.ConfirmPassword = confirm
	if r.Password != r.ConfirmPassword {
		return false
	}
	return true
}

func (r *RegisterInput) HashPassword() error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	r.Password = string(hashedPassword)

	return nil

}

func (r *RegisterInput) ValidateEmail() bool {

	cleaned := strings.TrimSpace(r.Email)
	cleaned = strings.ToLower(cleaned)
	r.Email = cleaned
	Pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(Pattern)
	return re.MatchString(r.Email)
}
